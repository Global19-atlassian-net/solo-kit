// Code generated by solo-kit. DO NOT EDIT.

package kubernetes

import (
	"sync"

	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/factory"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/wrapper"
	"github.com/solo-io/solo-kit/pkg/multicluster"
	"k8s.io/client-go/rest"
)

type PodMultiClusterClient interface {
	multicluster.ClusterHandler
	PodClient
}

type podMultiClusterClient struct {
	clients      map[string]PodClient
	clientAccess sync.RWMutex
	aggregator   wrapper.WatchAggregator
	cacheGetter  multicluster.KubeSharedCacheGetter
	opts         multicluster.KubeResourceFactoryOpts
}

func NewPodMultiClusterClient(cacheGetter multicluster.KubeSharedCacheGetter, opts multicluster.KubeResourceFactoryOpts) MockResourceMultiClusterClient {
	return NewMockResourceClientWithWatchAggregator(cacheGetter, nil, opts)
}

func NewPodMultiClusterClientWithWatchAggregator(cacheGetter multicluster.KubeSharedCacheGetter, aggregator wrapper.WatchAggregator, opts multicluster.KubeResourceFactoryOpts) MockResourceMultiClusterClient {
	return &podClientSet{
		clients:      make(map[string]PodClient),
		clientAccess: sync.RWMutex{},
		cacheGetter:  cacheGetter,
		aggregator:   aggregator,
		opts:         opts,
	}
}

func (c *podMultiClusterClient) clientFor(cluster string) (PodClient, error) {
	c.clientAccess.RLock()
	defer c.clientAccess.RUnlock()
	if client, ok := c.clients[cluster]; ok {
		return client, nil
	}
	return nil, multicluster.NoClientForClusterError(PodCrd.GroupVersionKind().String(), cluster)
}

func (c *podMultiClusterClient) ClusterAdded(cluster string, restConfig *rest.Config) {
	krc := &factory.KubeResourceClientFactory{
		Cluster:            cluster,
		Crd:                PodCrd,
		Cfg:                restConfig,
		SharedCache:        c.cacheGetter.GetCache(cluster),
		SkipCrdCreation:    c.opts.SkipCrdCreation,
		NamespaceWhitelist: c.opts.NamespaceWhitelist,
		ResyncPeriod:       c.opts.ResyncPeriod,
	}
	client, err := NewPodResourceClient(krc)
	if err != nil {
		return
	}
	if err := client.Register(); err != nil {
		return
	}
	c.clientAccess.Lock()
	defer c.clientAccess.Unlock()
	c.clients[cluster] = client
	if c.aggregator != nil {
		c.aggregator.AddWatch(client.BaseClient())
	}
}

func (c *podMultiClusterClient) ClusterRemoved(cluster string, restConfig *rest.Config) {
	c.clientAccess.Lock()
	defer c.clientAccess.Unlock()
	if client, ok := c.clients[cluster]; ok {
		delete(c.clients, cluster)
		if c.aggregator != nil {
			c.aggregator.RemoveWatch(client.BaseClient())
		}
	}
}

// TODO should we split this off the client interface?
func (c *podMultiClusterClient) BaseClient() clients.ResourceClient {
	panic("not implemented")
}

// TODO should we split this off the client interface?
func (c *podMultiClusterClient) Register() error {
	panic("not implemented")
}

func (c *podMultiClusterClient) Read(namespace, name string, opts clients.ReadOpts) (*Pod, error) {
	clusterClient, err := c.clientFor(opts.Cluster)
	if err != nil {
		return nil, err
	}
	return clusterClient.Read(namespace, name, opts)
}

func (c *podMultiClusterClient) Write(resource *Pod, opts clients.WriteOpts) (*Pod, error) {
	clusterClient, err := c.clientFor(resource.GetMetadata().GetCluster())
	if err != nil {
		return nil, err
	}
	return clusterClient.Write(resource, opts)
}

func (c *podMultiClusterClient) Delete(namespace, name string, opts clients.DeleteOpts) error {
	clusterClient, err := c.clientFor(opts.Cluster)
	if err != nil {
		return err
	}
	return clusterClient.Delete(namespace, name, opts)
}

func (c *podMultiClusterClient) List(namespace string, opts clients.ListOpts) (PodList, error) {
	clusterClient, err := c.clientFor(opts.Cluster)
	if err != nil {
		return nil, err
	}
	return clusterClient.List(namespace, opts)
}

func (c *podMultiClusterClient) Watch(namespace string, opts clients.WatchOpts) (<-chan PodList, <-chan error, error) {
	clusterClient, err := c.clientFor(opts.Cluster)
	if err != nil {
		return nil, nil, err
	}
	return clusterClient.Watch(namespace, opts)
}
