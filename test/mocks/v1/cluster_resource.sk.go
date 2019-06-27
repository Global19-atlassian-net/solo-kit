// Code generated by solo-kit. DO NOT EDIT.

package v1

import (
	"log"
	"sort"

	"github.com/solo-io/go-utils/hashutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/kube/crd"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"github.com/solo-io/solo-kit/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func NewClusterResource(namespace, name string) *ClusterResource {
	clusterresource := &ClusterResource{}
	clusterresource.SetMetadata(core.Metadata{
		Name:      name,
		Namespace: namespace,
	})
	return clusterresource
}

func (r *ClusterResource) SetMetadata(meta core.Metadata) {
	r.Metadata = meta
}

func (r *ClusterResource) SetStatus(status core.Status) {
	r.Status = status
}

func (r *ClusterResource) Hash() uint64 {
	metaCopy := r.GetMetadata()
	metaCopy.ResourceVersion = ""
	return hashutils.HashAll(
		metaCopy,
		r.BasicField,
	)
}

type ClusterResourceList []*ClusterResource

// namespace is optional, if left empty, names can collide if the list contains more than one with the same name
func (list ClusterResourceList) Find(namespace, name string) (*ClusterResource, error) {
	for _, clusterResource := range list {
		if clusterResource.GetMetadata().Name == name {
			if namespace == "" || clusterResource.GetMetadata().Namespace == namespace {
				return clusterResource, nil
			}
		}
	}
	return nil, errors.Errorf("list did not find clusterResource %v.%v", namespace, name)
}

func (list ClusterResourceList) AsResources() resources.ResourceList {
	var ress resources.ResourceList
	for _, clusterResource := range list {
		ress = append(ress, clusterResource)
	}
	return ress
}

func (list ClusterResourceList) AsInputResources() resources.InputResourceList {
	var ress resources.InputResourceList
	for _, clusterResource := range list {
		ress = append(ress, clusterResource)
	}
	return ress
}

func (list ClusterResourceList) Names() []string {
	var names []string
	for _, clusterResource := range list {
		names = append(names, clusterResource.GetMetadata().Name)
	}
	return names
}

func (list ClusterResourceList) NamespacesDotNames() []string {
	var names []string
	for _, clusterResource := range list {
		names = append(names, clusterResource.GetMetadata().Namespace+"."+clusterResource.GetMetadata().Name)
	}
	return names
}

func (list ClusterResourceList) Sort() ClusterResourceList {
	sort.SliceStable(list, func(i, j int) bool {
		return list[i].GetMetadata().Less(list[j].GetMetadata())
	})
	return list
}

func (list ClusterResourceList) Clone() ClusterResourceList {
	var clusterResourceList ClusterResourceList
	for _, clusterResource := range list {
		clusterResourceList = append(clusterResourceList, resources.Clone(clusterResource).(*ClusterResource))
	}
	return clusterResourceList
}

func (list ClusterResourceList) Each(f func(element *ClusterResource)) {
	for _, clusterResource := range list {
		f(clusterResource)
	}
}

func (list ClusterResourceList) EachResource(f func(element resources.Resource)) {
	for _, clusterResource := range list {
		f(clusterResource)
	}
}

func (list ClusterResourceList) AsInterfaces() []interface{} {
	var asInterfaces []interface{}
	list.Each(func(element *ClusterResource) {
		asInterfaces = append(asInterfaces, element)
	})
	return asInterfaces
}

var _ resources.Resource = &ClusterResource{}

// Kubernetes Adapter for ClusterResource

func (o *ClusterResource) GetObjectKind() schema.ObjectKind {
	t := ClusterResourceCrd.TypeMeta()
	return &t
}

func (o *ClusterResource) DeepCopyObject() runtime.Object {
	return resources.Clone(o).(*ClusterResource)
}

var (
	ClusterResourceCrd crd.Crd
	ClusterResourceGVK schema.GroupVersionKind
)

func init() {
	ClusterResourceGVK = schema.GroupVersionKind{
		Version: "v1",
		Group:   "crds.testing.solo.io",
		Kind:    "ClusterResource",
	}
	ClusterResourceCrd = crd.NewCrd(
		"clusterresources",
		ClusterResourceGVK.Group,
		ClusterResourceGVK.Version,
		ClusterResourceGVK.Kind,
		"clr",
		true,
		&ClusterResource{})
	if err := crd.GetRegistry().AddCrd(ClusterResourceCrd); err != nil {
		log.Fatalf("could not add crd to global registry")
	}
}
