// Code generated by solo-kit. DO NOT EDIT.

package v2alpha1

import (
	"sync"

	"github.com/solo-io/go-utils/errors"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/factory"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/wrapper"
	"github.com/solo-io/solo-kit/pkg/multicluster/handler"
	"k8s.io/client-go/rest"
)

var (
	NoMockResourceClientForClusterError = func(cluster string) error {
		return errors.Errorf("v2alpha1.MockResource client not found for cluster %v", cluster)
	}
)

type MockResourceMultiClusterClient interface {
	handler.ClusterHandler
	MockResourceInterface
}

type mockResourceMultiClusterClient struct {
	clients       map[string]MockResourceClient
	clientAccess  sync.RWMutex
	aggregator    wrapper.WatchAggregator
	factoryGetter factory.ResourceClientFactoryGetter
}

func NewMockResourceMultiClusterClient(factoryGetter factory.ResourceClientFactoryGetter) MockResourceMultiClusterClient {
	return NewMockResourceMultiClusterClientWithWatchAggregator(nil, factoryGetter)
}

func NewMockResourceMultiClusterClientWithWatchAggregator(aggregator wrapper.WatchAggregator, factoryGetter factory.ResourceClientFactoryGetter) MockResourceMultiClusterClient {
	return &mockResourceMultiClusterClient{
		clients:       make(map[string]MockResourceClient),
		clientAccess:  sync.RWMutex{},
		aggregator:    aggregator,
		factoryGetter: factoryGetter,
	}
}

func (c *mockResourceMultiClusterClient) interfaceFor(cluster string) (MockResourceInterface, error) {
	c.clientAccess.RLock()
	defer c.clientAccess.RUnlock()
	if client, ok := c.clients[cluster]; ok {
		return client, nil
	}
	return nil, NoMockResourceClientForClusterError(cluster)
}

func (c *mockResourceMultiClusterClient) ClusterAdded(cluster string, restConfig *rest.Config) {
	client, err := NewMockResourceClient(c.factoryGetter.ForCluster(cluster, restConfig))
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

func (c *mockResourceMultiClusterClient) ClusterRemoved(cluster string, restConfig *rest.Config) {
	c.clientAccess.Lock()
	defer c.clientAccess.Unlock()
	if client, ok := c.clients[cluster]; ok {
		delete(c.clients, cluster)
		if c.aggregator != nil {
			c.aggregator.RemoveWatch(client.BaseClient())
		}
	}
}

func (c *mockResourceMultiClusterClient) Read(namespace, name string, opts clients.ReadOpts) (*MockResource, error) {
	clusterInterface, err := c.interfaceFor(opts.Cluster)
	if err != nil {
		return nil, err
	}

	return clusterInterface.Read(namespace, name, opts)
}

func (c *mockResourceMultiClusterClient) Write(mockResource *MockResource, opts clients.WriteOpts) (*MockResource, error) {
	clusterInterface, err := c.interfaceFor(mockResource.GetMetadata().Cluster)
	if err != nil {
		return nil, err
	}
	return clusterInterface.Write(mockResource, opts)
}

func (c *mockResourceMultiClusterClient) Delete(namespace, name string, opts clients.DeleteOpts) error {
	clusterInterface, err := c.interfaceFor(opts.Cluster)
	if err != nil {
		return err
	}

	return clusterInterface.Delete(namespace, name, opts)
}

func (c *mockResourceMultiClusterClient) List(namespace string, opts clients.ListOpts) (MockResourceList, error) {
	clusterInterface, err := c.interfaceFor(opts.Cluster)
	if err != nil {
		return nil, err
	}

	return clusterInterface.List(namespace, opts)
}

func (c *mockResourceMultiClusterClient) Watch(namespace string, opts clients.WatchOpts) (<-chan MockResourceList, <-chan error, error) {
	clusterInterface, err := c.interfaceFor(opts.Cluster)
	if err != nil {
		return nil, nil, err
	}

	return clusterInterface.Watch(namespace, opts)
}
