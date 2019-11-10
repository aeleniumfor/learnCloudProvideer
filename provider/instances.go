package provider

import (
	"context"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	cloudprovider "k8s.io/cloud-provider"
)

type instances struct {
}

func newInstances() cloudprovider.Instances {
	return &instances{}
}

func (in instances) NodeAddresses(ctx context.Context, name types.NodeName) ([]v1.NodeAddress, error) {
	return nil, nil
}

func (in instances) NodeAddressesByProviderID(ctx context.Context, providerID string) ([]v1.NodeAddress, error) {
	return nil, nil
}

func (in instances) InstanceID(ctx context.Context, nodeName types.NodeName) (string, error) {
	return "", nil
}

func (in instances) InstanceType(ctx context.Context, name types.NodeName) (string, error) {
	return "", nil
}

func (in instances) InstanceTypeByProviderID(ctx context.Context, providerID string) (string, error) {
	return "", nil
}

func (in instances) AddSSHKeyToAllInstances(ctx context.Context, user string, keyData []byte) error {
	return nil
}

func (in instances) CurrentNodeName(ctx context.Context, hostname string) (types.NodeName, error) {
	return "", nil
}

func (in instances) InstanceExistsByProviderID(ctx context.Context, providerID string) (bool, error) {
	return true, nil
}

func (in instances) InstanceShutdownByProviderID(ctx context.Context, providerID string) (bool, error) {
	return true, nil
}
