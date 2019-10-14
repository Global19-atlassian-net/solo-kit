package main

import (
	"context"

	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/factory"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/wrapper"
	"github.com/solo-io/solo-kit/pkg/multicluster"
	"github.com/solo-io/solo-kit/pkg/multicluster/config"
	"github.com/solo-io/solo-kit/pkg/multicluster/handler"
	"github.com/solo-io/solo-kit/test/mocks/v2alpha1"
	"go.uber.org/zap"
)

func main() {
	ctx := context.Background()
	cm := multicluster.NewKubeSharedCacheManager(ctx)
	wa := wrapper.NewWatchAggregator()

	x := &factory.MultiClusterKubeResourceClientFactory{
		CacheGetter:        cm,
		WatchAggregator:    wa,
		Crd:                v2alpha1.MockResourceCrd,
		SkipCrdCreation:    false,
		NamespaceWhitelist: nil,
		ResyncPeriod:       0,
	}
	client, err := v2alpha1.NewMockResourceClient(x)
	if err != nil {
		contextutils.LoggerFrom(ctx).Fatalw("multi cluster client err!!!!!", zap.Error(err))
	}

	bc := client.BaseClient()
	ch, ok := bc.(handler.ClusterHandler)
	if !ok {
		contextutils.LoggerFrom(ctx).Fatalw("casting err!!!!!", zap.Error(err))
	}

	config.NewRestConfigHandler(multicluster.NewKubeConfigWatcher(), ch)

}
