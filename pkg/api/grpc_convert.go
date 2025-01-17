package api

import (
	"fmt"
	"vistara-node/pkg/api/types"
	"vistara-node/pkg/defaults"
	"vistara-node/pkg/models"
)

func convertMicroVMToModel(spec *types.MicroVMSpec) (*models.MicroVM, error) {
	vmid, err := models.NewVMID(spec.Id, defaults.MicroVMNamespace, "0")
	if err != nil {
		return nil, fmt.Errorf("creating vmid from spec: %w", err)
	}

	convertedModel := &models.MicroVM{
		ID: *vmid,
		Spec: models.MicroVMSpec{
			Kernel:     spec.KernelPath,
			VCPU:       spec.Vcpu,
			MemoryInMb: spec.MemoryInMb,
			HostNetDev: spec.HostNetDev,
			RootfsPath: spec.RootfsPath,
			GuestMAC:   spec.GuestMac,
		},
	}

	if convertedModel.Spec.VCPU == 0 {
		convertedModel.Spec.VCPU = defaults.VCPU
	}

	if convertedModel.Spec.MemoryInMb == 0 {
		convertedModel.Spec.MemoryInMb = defaults.MemoryInMb
	}

	return convertedModel, nil
}

func convertModelToMicroVMSpec(mvm *models.MicroVM) *types.MicroVMSpec {
	converted := &types.MicroVMSpec{
		Id:         mvm.ID.Name(),
		Vcpu:       mvm.Spec.VCPU,
		MemoryInMb: mvm.Spec.MemoryInMb,
		KernelPath: mvm.Spec.Kernel,
		RootfsPath: mvm.Spec.RootfsPath,
		HostNetDev: mvm.Spec.HostNetDev,
		GuestMac:   mvm.Spec.GuestMAC,
	}

	return converted
}

func String(s string) *string {
	return &s
}
