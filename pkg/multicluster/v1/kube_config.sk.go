// Code generated by solo-kit. DO NOT EDIT.

package v1

import (
	"sort"

	github_com_solo_io_solo_kit_api_multicluster_v1 "github.com/solo-io/solo-kit/api/multicluster/v1"

	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"github.com/solo-io/solo-kit/pkg/errors"
	"github.com/solo-io/solo-kit/pkg/utils/hashutils"
)

func NewKubeConfig(namespace, name string) *KubeConfig {
	kubeconfig := &KubeConfig{}
	kubeconfig.KubeConfig.SetMetadata(core.Metadata{
		Name:      name,
		Namespace: namespace,
	})
	return kubeconfig
}

// require custom resource to implement Clone() as well as resources.Resource interface

type CloneableKubeConfig interface {
	resources.Resource
	Clone() *github_com_solo_io_solo_kit_api_multicluster_v1.KubeConfig
}

var _ CloneableKubeConfig = &github_com_solo_io_solo_kit_api_multicluster_v1.KubeConfig{}

type KubeConfig struct {
	github_com_solo_io_solo_kit_api_multicluster_v1.KubeConfig
}

func (r *KubeConfig) Clone() resources.Resource {
	return &KubeConfig{KubeConfig: *r.KubeConfig.Clone()}
}

func (r *KubeConfig) Hash() uint64 {
	clone := r.KubeConfig.Clone()

	resources.UpdateMetadata(clone, func(meta *core.Metadata) {
		meta.ResourceVersion = ""
	})

	return hashutils.HashAll(clone)
}

type KubeConfigList []*KubeConfig
type KubeconfigsByNamespace map[string]KubeConfigList

// namespace is optional, if left empty, names can collide if the list contains more than one with the same name
func (list KubeConfigList) Find(namespace, name string) (*KubeConfig, error) {
	for _, kubeConfig := range list {
		if kubeConfig.GetMetadata().Name == name {
			if namespace == "" || kubeConfig.GetMetadata().Namespace == namespace {
				return kubeConfig, nil
			}
		}
	}
	return nil, errors.Errorf("list did not find kubeConfig %v.%v", namespace, name)
}

func (list KubeConfigList) AsResources() resources.ResourceList {
	var ress resources.ResourceList
	for _, kubeConfig := range list {
		ress = append(ress, kubeConfig)
	}
	return ress
}

func (list KubeConfigList) Names() []string {
	var names []string
	for _, kubeConfig := range list {
		names = append(names, kubeConfig.GetMetadata().Name)
	}
	return names
}

func (list KubeConfigList) NamespacesDotNames() []string {
	var names []string
	for _, kubeConfig := range list {
		names = append(names, kubeConfig.GetMetadata().Namespace+"."+kubeConfig.GetMetadata().Name)
	}
	return names
}

func (list KubeConfigList) Sort() KubeConfigList {
	sort.SliceStable(list, func(i, j int) bool {
		return list[i].GetMetadata().Less(list[j].GetMetadata())
	})
	return list
}

func (list KubeConfigList) Clone() KubeConfigList {
	var kubeConfigList KubeConfigList
	for _, kubeConfig := range list {
		kubeConfigList = append(kubeConfigList, resources.Clone(kubeConfig).(*KubeConfig))
	}
	return kubeConfigList
}

func (list KubeConfigList) Each(f func(element *KubeConfig)) {
	for _, kubeConfig := range list {
		f(kubeConfig)
	}
}

func (list KubeConfigList) AsInterfaces() []interface{} {
	var asInterfaces []interface{}
	list.Each(func(element *KubeConfig) {
		asInterfaces = append(asInterfaces, element)
	})
	return asInterfaces
}

func (byNamespace KubeconfigsByNamespace) Add(kubeConfig ...*KubeConfig) {
	for _, item := range kubeConfig {
		byNamespace[item.GetMetadata().Namespace] = append(byNamespace[item.GetMetadata().Namespace], item)
	}
}

func (byNamespace KubeconfigsByNamespace) Clear(namespace string) {
	delete(byNamespace, namespace)
}

func (byNamespace KubeconfigsByNamespace) List() KubeConfigList {
	var list KubeConfigList
	for _, kubeConfigList := range byNamespace {
		list = append(list, kubeConfigList...)
	}
	return list.Sort()
}

func (byNamespace KubeconfigsByNamespace) Clone() KubeconfigsByNamespace {
	cloned := make(KubeconfigsByNamespace)
	for ns, list := range byNamespace {
		cloned[ns] = list.Clone()
	}
	return cloned
}