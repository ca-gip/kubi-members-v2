// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	v1 "github.com/ca-gip/kubi-members/pkg/apis/cagip/v1"
	scheme "github.com/ca-gip/kubi-members/pkg/generated/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ClusterMembersGetter has a method to return a ClusterMemberInterface.
// A group's client should implement this interface.
type ClusterMembersGetter interface {
	ClusterMembers() ClusterMemberInterface
}

// ClusterMemberInterface has methods to work with ClusterMember resources.
type ClusterMemberInterface interface {
	Create(ctx context.Context, clusterMember *v1.ClusterMember, opts metav1.CreateOptions) (*v1.ClusterMember, error)
	Update(ctx context.Context, clusterMember *v1.ClusterMember, opts metav1.UpdateOptions) (*v1.ClusterMember, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.ClusterMember, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.ClusterMemberList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ClusterMember, err error)
	ClusterMemberExpansion
}

// clusterMembers implements ClusterMemberInterface
type clusterMembers struct {
	client rest.Interface
}

// newClusterMembers returns a ClusterMembers
func newClusterMembers(c *CagipV1Client) *clusterMembers {
	return &clusterMembers{
		client: c.RESTClient(),
	}
}

// Get takes name of the clusterMember, and returns the corresponding clusterMember object, and an error if there is any.
func (c *clusterMembers) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.ClusterMember, err error) {
	result = &v1.ClusterMember{}
	err = c.client.Get().
		Resource("clustermembers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ClusterMembers that match those selectors.
func (c *clusterMembers) List(ctx context.Context, opts metav1.ListOptions) (result *v1.ClusterMemberList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.ClusterMemberList{}
	err = c.client.Get().
		Resource("clustermembers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested clusterMembers.
func (c *clusterMembers) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("clustermembers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a clusterMember and creates it.  Returns the server's representation of the clusterMember, and an error, if there is any.
func (c *clusterMembers) Create(ctx context.Context, clusterMember *v1.ClusterMember, opts metav1.CreateOptions) (result *v1.ClusterMember, err error) {
	result = &v1.ClusterMember{}
	err = c.client.Post().
		Resource("clustermembers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterMember).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a clusterMember and updates it. Returns the server's representation of the clusterMember, and an error, if there is any.
func (c *clusterMembers) Update(ctx context.Context, clusterMember *v1.ClusterMember, opts metav1.UpdateOptions) (result *v1.ClusterMember, err error) {
	result = &v1.ClusterMember{}
	err = c.client.Put().
		Resource("clustermembers").
		Name(clusterMember.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterMember).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the clusterMember and deletes it. Returns an error if one occurs.
func (c *clusterMembers) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("clustermembers").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *clusterMembers) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("clustermembers").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched clusterMember.
func (c *clusterMembers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ClusterMember, err error) {
	result = &v1.ClusterMember{}
	err = c.client.Patch(pt).
		Resource("clustermembers").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
