package v1

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
	NoClusterResourceClientForClusterError = func(cluster string) error {
		return errors.Errorf("v1.ClusterResource client not found for cluster %v", cluster)
	}
)

type ClusterResourceMultiClusterClient interface {
	handler.ClusterHandler
	ClusterResourceInterface
}

type clusterResourceMultiClusterClient struct {
	clients       map[string]ClusterResourceClient
	clientAccess  sync.RWMutex
	aggregator    wrapper.WatchAggregator
	factoryGetter factory.ResourceClientFactoryGetter
}

func NewClusterResourceMultiClusterClient(factoryGetter factory.ResourceClientFactoryGetter) ClusterResourceMultiClusterClient {
	return NewClusterResourceMultiClusterClientWithWatchAggregator(nil, factoryGetter)
}

func NewClusterResourceMultiClusterClientWithWatchAggregator(aggregator wrapper.WatchAggregator, factoryGetter factory.ResourceClientFactoryGetter) ClusterResourceMultiClusterClient {
	return &clusterResourceMultiClusterClient{
		clients:       make(map[string]ClusterResourceClient),
		clientAccess:  sync.RWMutex{},
		aggregator:    aggregator,
		factoryGetter: factoryGetter,
	}
}

func (c *clusterResourceMultiClusterClient) interfaceFor(cluster string) (ClusterResourceInterface, error) {
	c.clientAccess.RLock()
	defer c.clientAccess.RUnlock()
	if client, ok := c.clients[cluster]; ok {
		return client, nil
	}
	return nil, NoClusterResourceClientForClusterError(cluster)
}

func (c *clusterResourceMultiClusterClient) ClusterAdded(cluster string, restConfig *rest.Config) {
	client, err := NewClusterResourceClient(c.factoryGetter.ForCluster(cluster, restConfig))
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

func (c *clusterResourceMultiClusterClient) ClusterRemoved(cluster string, restConfig *rest.Config) {
	c.clientAccess.Lock()
	defer c.clientAccess.Unlock()
	if client, ok := c.clients[cluster]; ok {
		delete(c.clients, cluster)
		if c.aggregator != nil {
			c.aggregator.RemoveWatch(client.BaseClient())
		}
	}
}

func (c *clusterResourceMultiClusterClient) Read(name string, opts clients.ReadOpts) (*ClusterResource, error) {
	clusterInterface, err := c.interfaceFor(opts.Cluster)
	if err != nil {
		return nil, err
	}

	return clusterInterface.Read(name, opts)
}

func (c *clusterResourceMultiClusterClient) Write(clusterResource *ClusterResource, opts clients.WriteOpts) (*ClusterResource, error) {
	clusterInterface, err := c.interfaceFor(clusterResource.GetMetadata().Cluster)
	if err != nil {
		return nil, err
	}
	return clusterInterface.Write(clusterResource, opts)
}

func (c *clusterResourceMultiClusterClient) Delete(name string, opts clients.DeleteOpts) error {
	clusterInterface, err := c.interfaceFor(opts.Cluster)
	if err != nil {
		return err
	}

	return clusterInterface.Delete(name, opts)
}

func (c *clusterResourceMultiClusterClient) List(opts clients.ListOpts) (ClusterResourceList, error) {
	clusterInterface, err := c.interfaceFor(opts.Cluster)
	if err != nil {
		return nil, err
	}

	return clusterInterface.List(opts)
}

func (c *clusterResourceMultiClusterClient) Watch(opts clients.WatchOpts) (<-chan ClusterResourceList, <-chan error, error) {
	clusterInterface, err := c.interfaceFor(opts.Cluster)
	if err != nil {
		return nil, nil, err
	}

	return clusterInterface.Watch(opts)
}
