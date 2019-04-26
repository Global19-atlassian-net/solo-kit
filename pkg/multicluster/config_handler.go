package multicluster

import (
	"context"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/kube/cache"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"sync"
)

type RestConfigs map[string]*rest.Config

type ClusterHandler interface {
	ClusterAdded(cluster string, restConfig *rest.Config)
	ClusterRemoved(cluster string, restConfig *rest.Config)
}

type RestConfigHandler struct {
	handlers []ClusterHandler
	cache    RestConfigs
	access   sync.RWMutex
}

func NewRestConfigHandler(handlers []ClusterHandler) *RestConfigHandler {
	return &RestConfigHandler{handlers: handlers}
}

func (h *RestConfigHandler) Run(ctx context.Context, local *rest.Config, kubeClient kubernetes.Interface, kubeCache cache.KubeCoreCache) error {
	kubeConfigs, errs, err := WatchKubeConfigs(ctx, kubeClient, kubeCache)
	if err != nil {
		return err
	}

	for {
		select {
		case err := <-errs:
			return err
		case kcs := <-kubeConfigs:
			restConfigs, err := makeRestConfigs(local, kcs)
			if err != nil {
				return err
			}

			h.handleNewRestConfigs(restConfigs)
		}
	}
}

func (h *RestConfigHandler) handleNewRestConfigs(cfgs RestConfigs) {
	h.access.RLock()
	for cluster, oldCfg := range h.cache {
		if _, persisted := cfgs[cluster]; persisted {
			continue
		}
		h.clusterRemoved(cluster, oldCfg)
	}
	for cluster, newCfg := range cfgs {
		if _, exists := h.cache[cluster]; exists {
			continue
		}
		h.clusterAdded(cluster, newCfg)
	}
	h.access.RUnlock()

	h.access.Lock()
	// update cache
	h.cache = cfgs
	h.access.Unlock()
}

func (h *RestConfigHandler) clusterAdded(cluster string, cfg *rest.Config) {
	h.access.RLock()
	defer h.access.RUnlock()
	for _, handler := range h.handlers {
		handler.ClusterAdded(cluster, cfg)
	}
}

func (h *RestConfigHandler) clusterRemoved(cluster string, cfg *rest.Config) {
	h.access.RLock()
	defer h.access.RUnlock()
	for _, handler := range h.handlers {
		handler.ClusterRemoved(cluster, cfg)
	}
}

func makeRestConfigs(local *rest.Config, kcs KubeConfigs) (RestConfigs, error) {
	cfgs := RestConfigs{LocalCluster: local}
	for cluster, kc := range kcs {
		restCfg, err := clientcmd.NewDefaultClientConfig(kc.Config, nil).ClientConfig()
		if err != nil {
			return nil, err
		}
		cfgs[cluster] = restCfg
	}
	return cfgs, nil
}
