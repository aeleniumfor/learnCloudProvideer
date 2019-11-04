package provider

import (
	"io"

	cloudprovider "k8s.io/cloud-provider"
)

type cloud struct {
	loadBalancers cloudprovider.LoadBalancer
}

// ProviderName is ProviderName
const ProviderName = "cloud"

func newCloud(configReader io.Reader) (cloudprovider.Interface, error) {
	return &cloud{
		loadBalancers: newLoadbalancers(),
	}, nil
}

func init() {
	cloudprovider.RegisterCloudProvider(ProviderName, func(config io.Reader) (cloudprovider.Interface, error) {
		return newCloud(config)
	})
}

func (c *cloud) Initialize(clientBuilder cloudprovider.ControllerClientBuilder, stop <-chan struct{}) {
}

func (c *cloud) LoadBalancer() (cloudprovider.LoadBalancer, bool) {
	return c.loadBalancers, true
}

func (c *cloud) Instances() (cloudprovider.Instances, bool) {
	return nil, true
}

func (c *cloud) Zones() (cloudprovider.Zones, bool) {
	return nil, true
}

func (c *cloud) Clusters() (cloudprovider.Clusters, bool) {
	return nil, false // not supported
}

func (c *cloud) Routes() (cloudprovider.Routes, bool) {
	return nil, false // not supported
}

func (c *cloud) ProviderName() string {
	return "cloud"
}

func (c *cloud) ScrubDNS(nameservers, searches []string) (nsOut, srchOut []string) {
	return nil, nil
}

func (c *cloud) HasClusterID() bool {
	return true
}
