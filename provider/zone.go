package provider

import (
	"context"

	"k8s.io/apimachinery/pkg/types"
	cloudprovider "k8s.io/cloud-provider"
)

type zones struct {
	region string
}

func newZones(region string) cloudprovider.Zones {
	return &zones{}
}

func (zn *zones) GetZone(ctx context.Context) (cloudprovider.Zone, error) {
	return cloudprovider.Zone{Region: zn.region}, nil
}

func (zn *zones) GetZoneByProviderID(ctx context.Context, providerID string) (cloudprovider.Zone, error) {
	return cloudprovider.Zone{Region: zn.region}, nil
}

func (zn *zones) GetZoneByNodeName(ctx context.Context, nodeName types.NodeName) (cloudprovider.Zone, error) {
	return cloudprovider.Zone{Region: zn.region}, nil
}
