package wrapper

import (
	"context"
	"sync"

	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
)

/*
A wrapper.watchAggregator wraps multiple ResourceWatchers and
aggregates a watch on each into a single Watch func
*/
type WatchAggregator interface {
	clients.ResourceWatcher
	AddWatch(w clients.ResourceWatcher) error
	RemoveWatch(w clients.ResourceWatcher)
}

type resourceSink chan resources.ResourceList

type addWatch func(watcher clients.ResourceWatcher) error
type removeWatch func()

type watchAggregator struct {
	sources        map[clients.ResourceWatcher][]removeWatch // how to unsubscribe this watcher
	sourcesAccess  sync.RWMutex
	sinks          map[resourceSink]addWatch // how to subscribe the aggregator to a watcher
	sinksAccess    sync.RWMutex
	watchers       map[clients.ResourceWatcher]struct{}
	watchersAccess sync.RWMutex
}

func NewWatchAggregator() WatchAggregator {
	sources := make(map[clients.ResourceWatcher][]removeWatch)
	sinks := make(map[resourceSink]addWatch)
	watchers := make(map[clients.ResourceWatcher]struct{})
	return &watchAggregator{sources: sources, sinks: sinks, watchers: watchers}
}

func (c *watchAggregator) Watch(namespace string, opts clients.WatchOpts) (<-chan resources.ResourceList, <-chan error, error) {
	opts = opts.WithDefaults()
	// create new sinks for this watch
	out := make(chan resources.ResourceList)
	aggregatedErrs := make(chan error)

	// a shared map that will be used to merge resources from different watchers
	listsByWatcher := make(resourcesByWatcher)
	access := sync.RWMutex{}

	// create a wait group for sources
	// so we can wait for all sources watches to close
	// before closing the sink channel (when this watch is canceled)
	sourceWatches := sync.WaitGroup{}

	// construct a func for adding an input watcher to this sink
	addWatch := func(watcher clients.ResourceWatcher) (err error) {
		sourceWatches.Add(1)

		// this function starts a watch for the watcher using the root context
		// we want to cancel it if :
		// the root context is cancelled
		// the watcher is removed
		ctx, cancel := context.WithCancel(opts.Ctx)

		// start a watch for the watcher on this namespace
		source, errs, err := watcher.Watch(namespace, clients.WatchOpts{
			Ctx:         ctx,
			RefreshRate: opts.RefreshRate,
			Selector:    opts.Selector,
		})
		if err != nil {
			return err
		}

		// read lists from the source channel,
		// group its resources by type
		go func() {
			defer cancel()
			defer sourceWatches.Done()

			for {
				select {
				case <-ctx.Done():
					return
				case err := <-errs:
					// if the source starts returning errors, remove its list from the snasphot
					access.Lock()
					delete(listsByWatcher, watcher)
					access.Unlock()
					aggregatedErrs <- err
					select {
					case <-ctx.Done():
						return
					case out <- listsByWatcher.merge():
					}
				case list, ok := <-source:
					if !ok {
						return
					}
					// add/update the list to the snapshot
					access.Lock()
					listsByWatcher[watcher] = list
					mergedList := listsByWatcher.merge()
					access.Unlock()
					select {
					case <-ctx.Done():
						return
					case out <- mergedList:
					}
				}
			}
		}()

		// construct a function for removing this watcher from this sink
		removeWatch := func() {
			// remove the watcher+resources from the snapshot
			access.Lock()
			delete(listsByWatcher, watcher)
			access.Unlock()
			cancel()
		}

		c.sourcesAccess.Lock()
		c.sources[watcher] = append(c.sources[watcher], removeWatch)
		c.sourcesAccess.Unlock()

		return nil
	}

	// add all the registered watchers to the sink
	c.watchersAccess.RLock()
	for w := range c.watchers {
		if err := addWatch(w); err != nil {
			return nil, nil, err
		}
	}
	c.watchersAccess.RUnlock()

	// store a way to add watches to this sink
	c.sinksAccess.Lock()
	c.sinks[out] = addWatch
	c.sinksAccess.Unlock()

	go func() {
		// context is closed, clean up watch resources
		<-opts.Ctx.Done()
		c.sinksAccess.Lock()
		delete(c.sinks, out)
		c.sinksAccess.Unlock()
		// wait for source watches to be closed before closing the sinks
		sourceWatches.Wait()
		close(out)
		close(aggregatedErrs)
	}()

	return out, aggregatedErrs, nil

}

func (c *watchAggregator) AddWatch(w clients.ResourceWatcher) error {
	c.watchersAccess.Lock()
	c.watchers[w] = struct{}{}
	c.watchersAccess.Unlock()

	c.sinksAccess.RLock()
	defer c.sinksAccess.RUnlock()
	for _, addWatcher := range c.sinks {
		if err := addWatcher(w); err != nil {
			return err
		}
	}
	return nil
}

func (c *watchAggregator) RemoveWatch(w clients.ResourceWatcher) {
	c.watchersAccess.Lock()
	delete(c.watchers, w)
	c.watchersAccess.Unlock()

	c.sourcesAccess.RLock()
	defer c.sourcesAccess.RUnlock()
	for _, removeWatcher := range c.sources[w] {
		removeWatcher()
	}
}

// aggregate resources by the channel they were read from
type resourcesByWatcher map[clients.ResourceWatcher]resources.ResourceList

func (rbw resourcesByWatcher) merge() resources.ResourceList {
	var merged resources.ResourceList
	for _, list := range rbw {
		merged = append(merged, list...)
	}
	return merged.Sort()
}
