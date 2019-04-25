package reconcile

import (
	"github.com/solo-io/go-utils/errors"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"go.uber.org/multierr"
)

/*
A MultiClusterReconciler takes a map of cluster names to resource clients
It can then perform reconciles against lists of desiredResources with resources
desired across multiple clusters
*/
type multiClusterReconciler struct {
	rcs map[string]clients.ResourceClient
}

func NewMultiClusterReconciler(rcs map[string]clients.ResourceClient) Reconciler {
	return &multiClusterReconciler{rcs: rcs}
}

func (r *multiClusterReconciler) Reconcile(namespace string, desiredResources resources.ResourceList, transitionFunc TransitionResourcesFunc, opts clients.ListOpts) error {
	byCluster := desiredResources.ByCluster()
	var errs error
	for cluster, desiredForCluster := range byCluster {
		rc, ok := r.rcs[cluster]
		if !ok {
			return errors.Errorf("no client found for cluster %v", cluster)
		}
		if err := NewReconciler(rc).Reconcile(namespace, desiredForCluster, transitionFunc, opts); err != nil {
			errs = multierr.Append(errs, errors.Wrapf(err, "reconciling cluster %v", cluster))
		}
	}
	return errs
}
