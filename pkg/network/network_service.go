package network

import (
	"context"
	ierror "errors"
	"fmt"
	"net"
	"strconv"
	"strings"
	"vistara-node/pkg/errors"
	"vistara-node/pkg/log"
	"vistara-node/pkg/models"
	"vistara-node/pkg/ports"

	"github.com/coreos/go-iptables/iptables"
	sysctl "github.com/lorenzosaino/go-sysctl"
	"github.com/sirupsen/logrus"
	"github.com/vishvananda/netlink"
)

type Config struct {
	ParentDeviceName string
	BridgeName       string
}

func New(cfg *Config) ports.NetworkService {
	return &networkService{
		parentDeviceName: cfg.ParentDeviceName,
		bridgeName:       cfg.BridgeName,
	}
}

type networkService struct {
	parentDeviceName string
	bridgeName       string
}

// IfaceCreate will create the network interface.
func (n *networkService) IfaceCreate(ctx context.Context, input ports.IfaceCreateInput) (*ports.IfaceDetails, error) {
	logger := log.GetLogger(ctx).WithFields(logrus.Fields{
		"service": "netlink_network",
		"iface":   input.DeviceName,
	})
	logger.Debugf(
		"creating network interface with type %s and MAC %s using parent %s",
		input.Type,
		input.MAC,
		n.parentDeviceName,
	)

	var (
		parentLink netlink.Link
		err        error
	)

	parentDeviceName := n.getParentIfaceName(input)
	if parentDeviceName == "" {
		return nil, errors.ErrParentIfaceRequiredForAttachingTap
	}

	var link netlink.Link

	switch input.Type {
	case models.IfaceTypeTap:
		link = &netlink.Tuntap{
			LinkAttrs: netlink.LinkAttrs{
				Name: input.DeviceName,
			},
			Mode: netlink.TuntapMode(netlink.TCA_CSUM_UPDATE_FLAG_ICMP),
		}
	case models.IfaceTypeMacvtap:
		link = &netlink.Macvtap{
			Macvlan: netlink.Macvlan{
				LinkAttrs: netlink.LinkAttrs{
					Name:        input.DeviceName,
					MTU:         parentLink.Attrs().MTU,
					ParentIndex: parentLink.Attrs().Index,
					Namespace:   parentLink.Attrs().Namespace,
					TxQLen:      parentLink.Attrs().TxQLen,
				},
				Mode: netlink.MACVLAN_MODE_BRIDGE,
			},
		}

		if input.MAC != "" {
			addr, parseErr := net.ParseMAC(input.MAC)
			if err != nil {
				return nil, fmt.Errorf("parsing mac address %s: %w", input.MAC, parseErr)
			}

			link.Attrs().HardwareAddr = addr
			logger.Tracef("added mac address %s to interface", addr)
		}
	case models.IfaceTypeUnsupported:
		return nil, errors.NewErrUnsupportedInterface(string(input.Type))
	default:
		return nil, errors.NewErrUnsupportedInterface(string(input.Type))
	}

	if err = netlink.LinkAdd(link); err != nil {
		return nil, fmt.Errorf("creating interface %s using netlink: %w", link.Attrs().Name, err)
	}

	macIf, err := netlink.LinkByName(link.Attrs().Name)
	if err != nil {
		return nil, fmt.Errorf("getting interface %s using netlink: %w", link.Attrs().Name, err)
	}

	if err := netlink.LinkSetUp(macIf); err != nil {
		return nil, fmt.Errorf("enabling device %s: %w", macIf.Attrs().Name, err)
	}

	logger.Debugf("created interface with mac %s", macIf.Attrs().HardwareAddr.String())

	if input.Type == models.IfaceTypeTap {
		deviceNameToIndex, err := strconv.Atoi(strings.ReplaceAll(input.DeviceName, "hypercore-", ""))
		if err != nil {
			return nil, fmt.Errorf("deviceName %s invalid", input.DeviceName)
		}

		addr, err := netlink.ParseAddr(GetTapDetails(deviceNameToIndex).TapIp.String() + "/30")
		if err != nil {
			return nil, fmt.Errorf("failed to parse TAP IP for index %d: %w", deviceNameToIndex, err)
		}

		if err = netlink.AddrAdd(link, addr); err != nil {
			return nil, fmt.Errorf("failed to add address to TAP device: %w", err)
		}

		err = sysctl.Set("net.ipv4.ip_forward", "1")
		if err != nil {
			return nil, fmt.Errorf("failed to enable IPv4 forwarding: %w", err)
		}

		err = sysctl.Set(fmt.Sprintf("net.ipv4.conf.%s.proxy_arp", link.Attrs().Name), "1")
		if err != nil {
			return nil, fmt.Errorf("failed to enable proxy_arp: %w", err)
		}

		err = sysctl.Set(fmt.Sprintf("net.ipv6.conf.%s.disable_ipv6", link.Attrs().Name), "1")
		if err != nil {
			return nil, fmt.Errorf("failed to enable disable_ipv6: %w", err)
		}

		ipt, err := iptables.New()
		if err != nil {
			return nil, fmt.Errorf("failed to create IPTables instance")
		}

		err = ipt.AppendUnique("nat", "POSTROUTING", "-o", parentDeviceName, "-j", "MASQUERADE")
		if err != nil {
			return nil, fmt.Errorf("failed to add MASQUERADE rule: %w", err)
		}

		err = ipt.InsertUnique("filter", "FORWARD", 1, "-m", "conntrack", "--ctstate", "RELATED,ESTABLISHED", "-j", "ACCEPT")
		if err != nil {
			return nil, fmt.Errorf("failed to add ACCEPT rule: %w", err)
		}

		err = ipt.InsertUnique("filter", "FORWARD", 1, "-i", link.Attrs().Name, "-o", parentDeviceName, "-j", "ACCEPT")
		if err != nil {
			return nil, fmt.Errorf("failed to add forwarding from parent device %s to TAP device %s: %w", parentDeviceName, link.Attrs().Name, err)
		}
	}

	return &ports.IfaceDetails{
		DeviceName: input.DeviceName,
		Type:       input.Type,
		MAC:        strings.ToUpper(macIf.Attrs().HardwareAddr.String()),
		Index:      macIf.Attrs().Index,
	}, nil
}

// IfaceDelete is used to delete a network interface.
func (n *networkService) IfaceDelete(ctx context.Context, input ports.DeleteIfaceInput) error {
	logger := log.GetLogger(ctx).WithFields(logrus.Fields{
		"service": "netlink_network",
		"iface":   input.DeviceName,
	})
	logger.Debug("deleting network interface")

	link, err := netlink.LinkByName(input.DeviceName)
	if err != nil {
		if ierror.Is(err, netlink.LinkNotFoundError{}) {
			return fmt.Errorf("failed to lookup network interface %s: %w", input.DeviceName, err)
		}

		logger.Debug("network interface doesn't exist, no action")

		return nil
	}

	if err = netlink.LinkDel(link); err != nil {
		return fmt.Errorf("deleting interface %s: %w", link.Attrs().Name, err)
	}

	return nil
}

func (n *networkService) IfaceExists(ctx context.Context, name string) (bool, error) {
	logger := log.GetLogger(ctx).WithFields(logrus.Fields{
		"service": "netlink_network",
		"iface":   name,
	})
	logger.Debug("checking if network interface exists")

	found, _, err := n.getIface(name)
	if err != nil {
		return false, fmt.Errorf("getting interface %s: %w", name, err)
	}

	return found, nil
}

// IfaceDetails will get the details of the supplied network interface.
func (n *networkService) IfaceDetails(ctx context.Context, name string) (*ports.IfaceDetails, error) {
	logger := log.GetLogger(ctx).WithFields(logrus.Fields{
		"service": "netlink_network",
		"iface":   name,
	})
	logger.Debug("getting network interface details")

	found, link, err := n.getIface(name)
	if err != nil {
		return nil, fmt.Errorf("getting interface %s: %w", name, err)
	}

	if !found {
		return nil, errors.ErrIfaceNotFound
	}

	details := &ports.IfaceDetails{
		DeviceName: name,
		MAC:        strings.ToUpper(link.Attrs().HardwareAddr.String()),
		Index:      link.Attrs().Index,
	}

	switch link.(type) {
	case *netlink.Macvtap:
		details.Type = models.IfaceTypeMacvtap
	case *netlink.Tuntap:
		details.Type = models.IfaceTypeTap
	default:
		details.Type = models.IfaceTypeUnsupported
	}

	return details, nil
}

func (n *networkService) getIface(name string) (bool, netlink.Link, error) {
	link, err := netlink.LinkByName(name)
	if err != nil {
		if ierror.Is(err, netlink.LinkNotFoundError{}) {
			return false, nil, fmt.Errorf("failed to lookup network interface %s: %w", name, err)
		}

		return false, nil, nil
	}

	return true, link, nil
}

func (n *networkService) getParentIfaceName(input ports.IfaceCreateInput) string {
	if input.Type == models.IfaceTypeMacvtap {
		return n.parentDeviceName
	}

	if input.BridgeName != "" {
		return input.BridgeName
	}

	return n.bridgeName
}
