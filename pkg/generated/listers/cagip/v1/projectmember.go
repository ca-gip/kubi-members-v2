// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/ca-gip/kubi-members-v2/pkg/apis/cagip/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ProjectMemberLister helps list ProjectMembers.
// All objects returned here must be treated as read-only.
type ProjectMemberLister interface {
	// List lists all ProjectMembers in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.ProjectMember, err error)
	// ProjectMembers returns an object that can list and get ProjectMembers.
	ProjectMembers(namespace string) ProjectMemberNamespaceLister
	ProjectMemberListerExpansion
}

// projectMemberLister implements the ProjectMemberLister interface.
type projectMemberLister struct {
	indexer cache.Indexer
}

// NewProjectMemberLister returns a new ProjectMemberLister.
func NewProjectMemberLister(indexer cache.Indexer) ProjectMemberLister {
	return &projectMemberLister{indexer: indexer}
}

// List lists all ProjectMembers in the indexer.
func (s *projectMemberLister) List(selector labels.Selector) (ret []*v1.ProjectMember, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ProjectMember))
	})
	return ret, err
}

// ProjectMembers returns an object that can list and get ProjectMembers.
func (s *projectMemberLister) ProjectMembers(namespace string) ProjectMemberNamespaceLister {
	return projectMemberNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ProjectMemberNamespaceLister helps list and get ProjectMembers.
// All objects returned here must be treated as read-only.
type ProjectMemberNamespaceLister interface {
	// List lists all ProjectMembers in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.ProjectMember, err error)
	// Get retrieves the ProjectMember from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.ProjectMember, error)
	ProjectMemberNamespaceListerExpansion
}

// projectMemberNamespaceLister implements the ProjectMemberNamespaceLister
// interface.
type projectMemberNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ProjectMembers in the indexer for a given namespace.
func (s projectMemberNamespaceLister) List(selector labels.Selector) (ret []*v1.ProjectMember, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ProjectMember))
	})
	return ret, err
}

// Get retrieves the ProjectMember from the indexer for a given namespace and name.
func (s projectMemberNamespaceLister) Get(name string) (*v1.ProjectMember, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("projectmember"), name)
	}
	return obj.(*v1.ProjectMember), nil
}
