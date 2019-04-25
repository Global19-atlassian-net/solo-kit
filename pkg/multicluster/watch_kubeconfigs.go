package multicluster

import (
	"context"
	"github.com/solo-io/go-utils/contextutils"
	v1 "github.com/solo-io/solo-kit/pkg/multicluster/v1"

	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/factory"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/kube/cache"
	"github.com/solo-io/solo-kit/pkg/multicluster/secretconverter"
	"k8s.io/client-go/kubernetes"
)

// empty ClusterId refers to local
const LocalCluster = ""

type KubeConfigs map[string]*v1.KubeConfig

func WatchKubeConfigs(ctx context.Context, kube kubernetes.Interface, cache cache.KubeCoreCache) (<-chan KubeConfigs, <-chan error, error) {
	kubeConfigClient, err := v1.NewKubeConfigClient(&factory.KubeSecretClientFactory{
		Clientset:       kube,
		Cache:           cache,
		SecretConverter: &secretconverter.KubeConfigSecretConverter{},
	})
	if err != nil {
		return nil, nil, err
	}
	emitter := v1.NewKubeconfigsEmitter(kubeConfigClient)
	kubeConfigsChan := make(chan KubeConfigs)
	eventLoop := v1.NewKubeconfigsEventLoop(emitter, &kubeConfigSyncer{kubeConfigsChan: kubeConfigsChan})
	errs, err := eventLoop.Run(nil, clients.WatchOpts{Ctx: ctx})
	if err != nil {
		return nil, nil, err
	}
	return kubeConfigsChan, errs, nil
}

type kubeConfigSyncer struct {
	kubeConfigsChan chan KubeConfigs
}

func (s *kubeConfigSyncer) Sync(ctx context.Context, snap *v1.KubeconfigsSnapshot) error {
	ctx = contextutils.WithLogger(ctx, "multicluster")
	logger := contextutils.LoggerFrom(ctx)
	cfgs := snap.Kubeconfigs.List()
	kubeConfigs := make(KubeConfigs)
	for _, cfg := range cfgs {
		if _, alreadyDefined := kubeConfigs[cfg.Cluster]; alreadyDefined {
			logger.Warnf("secret already defined for %v, %v will be ignored", cfg.Cluster, cfg.Metadata.Ref())
			continue
		}
		kubeConfigs[cfg.Cluster] = cfg
	}
	s.kubeConfigsChan <- kubeConfigs
	return nil
}
