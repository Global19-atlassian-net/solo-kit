package wrapper

import (
	"sync"

	"github.com/solo-io/go-utils/errutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
)

/*
A wrapper.MultiWatcher wraps multiple ResourceWatchers and
aggregates a watch on each into a single Watch func
*/
type MultiWatcher struct {
	Watchers []clients.ResourceWatcher
}

func (c *MultiWatcher) Watch(namespace string, opts clients.WatchOpts) (<-chan resources.ResourceList, <-chan error, error) {
	opts = opts.WithDefaults()
	aggregatedErrs := make(chan error)

	access := sync.Mutex{}
	resourcesByWatcher := make(map[clients.ResourceWatcher]resources.ResourceList)

	aggregatedOut := make(chan resources.ResourceList)

	wg := sync.WaitGroup{}
	for _, watcher := range c.Watchers {
		watcher := watcher
		resourceLists, errs, err := watcher.Watch(namespace, opts)
		if err != nil {
			return nil, nil, err
		}
		wg.Add(1)
		go errutils.AggregateErrs(opts.Ctx, aggregatedErrs, errs, "multiwatch")
		go func() {
			defer wg.Done()
			for list := range resourceLists {
				access.Lock()
				resourcesByWatcher[watcher] = list
				var aggregatedResources resources.ResourceList
				for _, list := range resourcesByWatcher {
					aggregatedResources = append(aggregatedResources, list.Copy()...)
				}
				aggregatedResources = aggregatedResources.Sort()
				access.Unlock()
				select {
				case aggregatedOut <- aggregatedResources:
				case <-opts.Ctx.Done():
					return
				}
			}
		}()
	}
	go func() {
		wg.Wait()
		close(aggregatedOut)
	}()
	return aggregatedOut, aggregatedErrs, nil

}
