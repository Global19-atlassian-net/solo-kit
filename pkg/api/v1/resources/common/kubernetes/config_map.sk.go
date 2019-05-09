// Code generated by solo-kit. DO NOT EDIT.

package kubernetes

import (
	"sort"

	github_com_solo_io_solo_kit_api_external_kubernetes_configmap "github.com/solo-io/solo-kit/api/external/kubernetes/configmap"

	"github.com/solo-io/go-utils/hashutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"github.com/solo-io/solo-kit/pkg/errors"
)

func NewConfigMap(namespace, name string) *ConfigMap {
	configmap := &ConfigMap{}
	configmap.ConfigMap.SetMetadata(core.Metadata{
		Name:      name,
		Namespace: namespace,
	})
	return configmap
}

// require custom resource to implement Clone() as well as resources.Resource interface

type CloneableConfigMap interface {
	resources.Resource
	Clone() *github_com_solo_io_solo_kit_api_external_kubernetes_configmap.ConfigMap
}

var _ CloneableConfigMap = &github_com_solo_io_solo_kit_api_external_kubernetes_configmap.ConfigMap{}

type ConfigMap struct {
	github_com_solo_io_solo_kit_api_external_kubernetes_configmap.ConfigMap
}

func (r *ConfigMap) Clone() resources.Resource {
	return &ConfigMap{ConfigMap: *r.ConfigMap.Clone()}
}

func (r *ConfigMap) Hash() uint64 {
	clone := r.ConfigMap.Clone()

	resources.UpdateMetadata(clone, func(meta *core.Metadata) {
		meta.ResourceVersion = ""
	})

	return hashutils.HashAll(clone)
}

type ConfigMapList []*ConfigMap

// namespace is optional, if left empty, names can collide if the list contains more than one with the same name
func (list ConfigMapList) Find(namespace, name string) (*ConfigMap, error) {
	for _, configMap := range list {
		if configMap.GetMetadata().Name == name {
			if namespace == "" || configMap.GetMetadata().Namespace == namespace {
				return configMap, nil
			}
		}
	}
	return nil, errors.Errorf("list did not find configMap %v.%v", namespace, name)
}

func (list ConfigMapList) AsResources() resources.ResourceList {
	var ress resources.ResourceList
	for _, configMap := range list {
		ress = append(ress, configMap)
	}
	return ress
}

func (list ConfigMapList) Names() []string {
	var names []string
	for _, configMap := range list {
		names = append(names, configMap.GetMetadata().Name)
	}
	return names
}

func (list ConfigMapList) NamespacesDotNames() []string {
	var names []string
	for _, configMap := range list {
		names = append(names, configMap.GetMetadata().Namespace+"."+configMap.GetMetadata().Name)
	}
	return names
}

func (list ConfigMapList) Sort() ConfigMapList {
	sort.SliceStable(list, func(i, j int) bool {
		return list[i].GetMetadata().Less(list[j].GetMetadata())
	})
	return list
}

func (list ConfigMapList) Clone() ConfigMapList {
	var configMapList ConfigMapList
	for _, configMap := range list {
		configMapList = append(configMapList, resources.Clone(configMap).(*ConfigMap))
	}
	return configMapList
}

func (list ConfigMapList) Each(f func(element *ConfigMap)) {
	for _, configMap := range list {
		f(configMap)
	}
}

func (list ConfigMapList) EachResource(f func(element resources.Resource)) {
	for _, configMap := range list {
		f(configMap)
	}
}

func (list ConfigMapList) AsInterfaces() []interface{} {
	var asInterfaces []interface{}
	list.Each(func(element *ConfigMap) {
		asInterfaces = append(asInterfaces, element)
	})
	return asInterfaces
}
