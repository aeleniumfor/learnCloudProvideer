package provider

import (
	"fmt"
	"io"
	"os"

	cloudprovider "k8s.io/cloud-provider"
)

type cloud struct {
	loadBalancers cloudprovider.LoadBalancer
	// instances     cloudprovider.Instances
	// zones         cloudprovider.Zones
}

const (
	//ProviderName is Cloud-Provider-Name
	ProviderName = "testcloud"

	// TagsKubernetesResource is a marker tag indicating that resource is part of k8s cluster
	TagsKubernetesResource = "@k8s"

	// TagsClusterID is tag name for mark ClusterID
	TagsClusterID = TagsKubernetesResource + ".ClusterID"
)

func newCloud() (cloudprovider.Interface, error) {
	fmt.Fprintf(os.Stderr, "star: %v\n", "newCloud")
	return &cloud{
		loadBalancers: newLoadbalancers(),
		// instances:     newInstances(),
		// zones:         newZones("testcloud"),
	}, nil
}

func init() {
	cloudprovider.RegisterCloudProvider(ProviderName, func(io.Reader) (cloudprovider.Interface, error) {
		fmt.Println(ProviderName)
		return newCloud()
	})
}

func (c *cloud) Initialize(clientBuilder cloudprovider.ControllerClientBuilder, stop <-chan struct{}) {
	fmt.Println("start cloud.go Initialize")
}

func (c *cloud) LoadBalancer() (cloudprovider.LoadBalancer, bool) {
	fmt.Println("loadBarancer")
	fmt.Println("loadBarancer")
	fmt.Println("loadBarancer")

	return c.loadBalancers, true
}

func (c *cloud) Instances() (cloudprovider.Instances, bool) {
	return nil, false
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
