package provider

import (
	"context"

	cloudprovider "k8s.io/cloud-provider"
)

type clusters struct {
}

func newClusters() cloudprovider.Clusters {
	return &clusters{}
}

func (c *clusters) ListClusters(ctx context.Context) ([]string, error) {
	return nil, nil
}

func (c *clusters) Master(ctx context.Context, clusterName string) (string, error) {
	return "", nil
}
