package provider

import (
	"context"
	"fmt"

	v1 "k8s.io/api/core/v1"
	cloudprovider "k8s.io/cloud-provider"
)

type loadbalancers struct {
}

// newLoadbalancers returns a cloudprovider.LoadBalancer whose concrete type is a *loadbalancer.
func newLoadbalancers() cloudprovider.LoadBalancer {
	return &loadbalancers{}
}

func (lb *loadbalancers) GetLoadBalancer(ctx context.Context, clusterName string, service *v1.Service) (status *v1.LoadBalancerStatus, exists bool, err error) {
	// GetLoadBalancer(ctx context.Context, clusterName string, service *v1.Service) (status *v1.LoadBalancerStatus, exists bool, err error)
	lb.GetLoadBalancerName(ctx, clusterName, service)
	// status = &v1.LoadBalancerStatus{}
	// status.Ingress = append(status.Ingress, v1.LoadBalancerIngress{IP: "192.168.56.104"})

	return nil, true, nil
}

func (lb *loadbalancers) GetLoadBalancerName(ctx context.Context, clusterName string, service *v1.Service) string {
	// GetLoadBalancerName(ctx context.Context, clusterName string, service *v1.Service) string
	return cloudprovider.DefaultLoadBalancerName(service)
}

func (lb *loadbalancers) EnsureLoadBalancer(ctx context.Context, clusterName string, service *v1.Service, nodes []*v1.Node) (*v1.LoadBalancerStatus, error) {
	// EnsureLoadBalancer(ctx context.Context, clusterName string, service *v1.Service, nodes []*v1.Node) (*v1.LoadBalancerStatus, error)
	_, exists, err := lb.GetLoadBalancer(ctx, clusterName, service)
	if err != nil {
		return nil, err
	}
	fmt.Println(exists)
	lbStatus := &v1.LoadBalancerStatus{}
	lbStatus.Ingress = []v1.LoadBalancerIngress{{IP: "192.168.11.01"}}
	return lbStatus, nil

}

func (lb *loadbalancers) UpdateLoadBalancer(ctx context.Context, clusterName string, service *v1.Service, nodes []*v1.Node) error {
	// UpdateLoadBalancer(ctx context.Context, clusterName string, service *v1.Service, nodes []*v1.Node) error
	return nil
}
func (lb *loadbalancers) EnsureLoadBalancerDeleted(ctx context.Context, clusterName string, service *v1.Service) error {
	// EnsureLoadBalancerDeleted(ctx context.Context, clusterName string, service *v1.Service) error
	return nil
}
