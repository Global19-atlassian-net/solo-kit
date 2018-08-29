package pkg

import (
	"context"
	"errors"
	"net/url"
	"sync"

	"github.com/solo-io/solo-kit/pkg/utils/contextutils"
	"go.uber.org/zap"

	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/projects/gloo/pkg/api/v1"
	"github.com/solo-io/solo-kit/projects/gloo/pkg/api/v1/plugins"
)

type Updater struct {
	functionalPlugins []FuncitonDiscovery
	activeupstreams   map[string]context.CancelFunc
	ctx               context.Context
	resolver          Resolver
	logger            *zap.SugaredLogger

	upstreamClient v1.UpstreamClient
	secretClient   v1.SecretClient

	maxInParallelSemaphore chan struct{}
}

func getConcurrencyChan(maxoncurrency uint) chan struct{} {
	if maxoncurrency == 0 {
		return nil
	}
	ret := make(chan struct{}, maxoncurrency)
	go func() {
		for i := uint(0); i < maxoncurrency; i++ {
			ret <- struct{}{}
		}
	}()
	return ret

}

func New(ctx context.Context, resolver Resolver, maxoncurrency uint, functionalPlugins []FuncitonDiscovery) *Updater {
	ctx = contextutils.WithLogger(ctx, "function-discovery-updater")
	return &Updater{
		logger:                 contextutils.LoggerFrom(ctx),
		ctx:                    ctx,
		resolver:               resolver,
		functionalPlugins:      functionalPlugins,
		activeupstreams:        make(map[string]context.CancelFunc),
		maxInParallelSemaphore: getConcurrencyChan(maxoncurrency),
	}
}

type detectResult struct {
	spec *plugins.ServiceSpec
	fp   FuncitonDiscovery
}

func (u *Updater) detectType(ctx context.Context, url *url.URL) (*detectResult, error) {
	// TODO add global timeout?
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	result := make(chan detectResult, 1)

	// run all detections in paralel
	var waitgroup sync.WaitGroup
	for _, fp := range u.functionalPlugins {
		waitgroup.Add(1)
		go func(functionalPlugins FuncitonDiscovery) {
			defer waitgroup.Done()

			if u.maxInParallelSemaphore != nil {
				select {
				case token := <-u.maxInParallelSemaphore:
					defer func() { u.maxInParallelSemaphore <- token }()
				case <-ctx.Done():
					return //ctx.Err()
				}
			}

			spec, err := functionalPlugins.DetectUpstreamType(ctx, url)
			if err == nil && spec != nil {
				// success
				result <- detectResult{
					spec: spec,
					fp:   functionalPlugins,
				}
			}
			if err == context.Canceled {
				return
			}
			if err != nil {
				// TODO retry + backoff:(
			}
		}(fp)
	}
	go func() {
		waitgroup.Wait()
		close(result)

	}()
	select {
	case res, ok := <-result:
		if ok {
			return &res, nil
		}
		return nil, errors.New("upstream type cannot be detected")
	case <-ctx.Done():
		return nil, ctx.Err()
	}

}

func (u *Updater) saveUpstream(ctx context.Context, upstream *v1.Upstream, mutator UpstreamMutator) error {
	err := mutator(upstream)
	if err != nil {
		return err
	}

	var wo clients.WriteOpts
	wo.Ctx = ctx
	wo.OverwriteExisting = true

	/* upstream, err = */
	u.upstreamClient.Write(upstream, wo)

	// TODO: if write failed, due to resource conflict,
	// get latest version, and if it still doesnt have a spec, mutate again and retry.

	return nil
}

type supportSpec interface {
	SetServiceSpec(*plugins.ServiceSpec)
	GetServiceSpec()
}

func (u *Updater) Run() error {

	// watch upstreams and the such.
	return nil

}

func (u *Updater) UpstreamAdded(upstream *v1.Upstream) {
	// upstream already tracked. ignore.
	key := resources.Key(upstream)
	if _, ok := u.activeupstreams[key]; ok {
		return
	}
	ctx, cancel := context.WithCancel(u.ctx)
	go u.RunForUpstream(ctx, upstream)
	u.activeupstreams[key] = cancel
}

func (u *Updater) UpstreamRemoved(upstream *v1.Upstream) {
	key := resources.Key(upstream)
	if cancel, ok := u.activeupstreams[key]; ok {
		cancel()
		delete(u.activeupstreams, key)
	}
}

func (u *Updater) RunForUpstream(ctx context.Context, upstream *v1.Upstream) error {

	// see if anyone likes this upstream:
	var discoveryForUpstream FuncitonDiscovery
	for _, fp := range u.functionalPlugins {
		if fp.IsUpstreamFunctional(upstream) {
			discoveryForUpstream = fp
			break
		}
	}

	upstreamSave := func(m UpstreamMutator) error {
		return u.saveUpstream(ctx, upstream, m)
	}

	if discoveryForUpstream == nil {
		// TODO: this will probably not going to work unless the upstream type will also have the method required
		_, ok := upstream.UpstreamSpec.UpstreamType.(supportSpec)
		if !ok {
			// can't set a service spec - which is required from this point on, as hueristic detection requires spec
			return errors.New("discovery not possible for upsteram")
		}

		// if we are here it means that the service upstream doesn't have a spec
		url, err := u.resolver.Resolve(upstream)
		if err != nil {
			return err
		}
		// try to detect the type
		res, err := u.detectType(ctx, url)
		if err != nil {
			return err
		}
		discoveryForUpstream = res.fp
		upstreamSave(func(upstream *v1.Upstream) error {

			servicespecupstream, ok := upstream.UpstreamSpec.UpstreamType.(supportSpec)
			if !ok {
				return errors.New("can't set spec")
			}
			servicespecupstream.SetServiceSpec(res.spec)
			return nil
		})

	}

	// TODO: figure out how to get the secret list
	return discoveryForUpstream.DetectFunctions(ctx, nil, upstream, upstreamSave)
}