package provider

import (
	"fmt"
	"io"
	"os"

	cloudprovider "k8s.io/cloud-provider"
)

type cloud struct {
	loadBalancers cloudprovider.LoadBalancer
	instances     cloudprovider.Instances
}

// ProviderName is ProviderName
const ProviderName = "testcloud"

func newCloud() (cloudprovider.Interface, error) {
	fmt.Fprintf(os.Stderr, "star: %v\n", "newCloud")
	return &cloud{
		loadBalancers: newLoadbalancers(),
		instances:     newInstances(),
	}, nil
}

func init() {
	cloudprovider.RegisterCloudProvider(ProviderName, func(io.Reader) (cloudprovider.Interface, error) {
		return newCloud()
	})
}

func (c *cloud) Initialize(clientBuilder cloudprovider.ControllerClientBuilder, stop <-chan struct{}) {
}

func (c *cloud) LoadBalancer() (cloudprovider.LoadBalancer, bool) {
	return c.loadBalancers, true
}

func (c *cloud) Instances() (cloudprovider.Instances, bool) {
	return c.instances, true
}

func (c *cloud) Zones() (cloudprovider.Zones, bool) {
	return nil, false
}

func (c *cloud) Clusters() (cloudprovider.Clusters, bool) {
	return nil, false // not supported
}

func (c *cloud) Routes() (cloudprovider.Routes, bool) {
	return nil, false // not supported
}

func (c *cloud) ProviderName() string {
	return ProviderName
}

func (c *cloud) ScrubDNS(nameservers, searches []string) (nsOut, srchOut []string) {
	return nil, nil
}

func (c *cloud) HasClusterID() bool {
	return true
}
