/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by informer-gen. DO NOT EDIT.

package v2alpha1

import (
	time "time"

	testingsoloiov2alpha1 "github.com/solo-io/solo-kit/test/mocks/v2alpha1/kube/apis/testing.solo.io/v2alpha1"
	versioned "github.com/solo-io/solo-kit/test/mocks/v2alpha1/kube/client/clientset/versioned"
	internalinterfaces "github.com/solo-io/solo-kit/test/mocks/v2alpha1/kube/client/informers/externalversions/internalinterfaces"
	v2alpha1 "github.com/solo-io/solo-kit/test/mocks/v2alpha1/kube/client/listers/testing.solo.io/v2alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// MockResourceInformer provides access to a shared informer and lister for
// MockResources.
type MockResourceInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v2alpha1.MockResourceLister
}

type mockResourceInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewMockResourceInformer constructs a new informer for MockResource type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewMockResourceInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredMockResourceInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredMockResourceInformer constructs a new informer for MockResource type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredMockResourceInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TestingV2alpha1().MockResources(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TestingV2alpha1().MockResources(namespace).Watch(options)
			},
		},
		&testingsoloiov2alpha1.MockResource{},
		resyncPeriod,
		indexers,
	)
}

func (f *mockResourceInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredMockResourceInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *mockResourceInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&testingsoloiov2alpha1.MockResource{}, f.defaultInformer)
}

func (f *mockResourceInformer) Lister() v2alpha1.MockResourceLister {
	return v2alpha1.NewMockResourceLister(f.Informer().GetIndexer())
}