// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "github.com/ca-gip/kubi-members/pkg/generated/clientset/versioned/typed/ca-gip/v1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeCagipV1 struct {
	*testing.Fake
}

func (c *FakeCagipV1) ClusterMembers() v1.ClusterMemberInterface {
	return &FakeClusterMembers{c}
}

func (c *FakeCagipV1) ProjectMembers(namespace string) v1.ProjectMemberInterface {
	return &FakeProjectMembers{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeCagipV1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
