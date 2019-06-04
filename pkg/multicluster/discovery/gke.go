package discovery

import (
	"context"
	"encoding/base64"
	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/go-utils/errors"
	"github.com/solo-io/go-utils/kubeutils"
	apiv1 "github.com/solo-io/solo-kit/api/multicluster/v1"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	v1 "github.com/solo-io/solo-kit/pkg/multicluster/v1"
	"google.golang.org/api/container/v1"
	"google.golang.org/api/option"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/clientcmd/api"
	"net/http"
	"net/url"
)

type ClusterDiscoveryOpts struct {
	ProjectID            string
	Zone                 string
	ClusterLabelSelector string
	CredentialsJSON      []byte
}

func DiscoverGkeClusters(ctx context.Context, writeNamespace string, opts ClusterDiscoveryOpts) (v1.KubeConfigList, error) {
	ctx = contextutils.WithLogger(ctx, "gke-cluster-discovery")
	gke, err := container.NewService(ctx, option.WithCredentialsJSON(opts.CredentialsJSON))
	if err != nil {
		return nil, err
	}
	zone := opts.Zone
	if zone == "" {
		zone = "-"
	}
	listResponse, err := gke.Projects.Zones.Clusters.List(opts.ProjectID, zone).Do()
	if err != nil {
		return nil, err
	}
	if listResponse.HTTPStatusCode != http.StatusOK {
		return nil, errors.Errorf("non-200 status code: %v", listResponse.ServerResponse)
	}

	var labelSelector labels.Selector
	if opts.ClusterLabelSelector != "" {
		labelSelector, err = labels.Parse(opts.ClusterLabelSelector)
		if err != nil {
			return nil, err
		}
	} else {
		labelSelector = labels.Everything()
	}

	var kcs v1.KubeConfigList
	for _, cluster := range listResponse.Clusters {
		if !labelSelector.Matches(labels.Set(cluster.ResourceLabels)) {
			continue
		}
		clusterName := cluster.Name
		if cluster.MasterAuth == nil {
			contextutils.LoggerFrom(ctx).Errorf("invalid response from GKE: cluster %v missing masterAuth field on cluster",
				clusterName)
			continue
		}
		server := &url.URL{Scheme: "https", Host: cluster.Endpoint}
		encodedClientCert := cluster.MasterAuth.ClientCertificate
		if encodedClientCert == "" {
			encodedClientCert = cluster.MasterAuth.ClusterCaCertificate
		}
		clientCert, err := base64.StdEncoding.DecodeString(encodedClientCert)
		if err != nil {
			return nil, err
		}
		kcs = append(kcs, &v1.KubeConfig{
			KubeConfig: apiv1.KubeConfig{
				Metadata: core.Metadata{
					Namespace: writeNamespace,
					Name:      kubeutils.SanitizeName(clusterName),
				},
				Cluster: clusterName,
				Config: api.Config{
					CurrentContext: clusterName,
					Contexts: map[string]*api.Context{
						clusterName: {
							AuthInfo: clusterName,
							Cluster:  clusterName,
						},
					},
					Clusters: map[string]*api.Cluster{
						clusterName: {
							Server:                   server.String(),
							CertificateAuthorityData: clientCert,
						},
					},
					AuthInfos: map[string]*api.AuthInfo{
						clusterName: {
							AuthProvider: &api.AuthProviderConfig{
								Name:"gcp",
								Config: map[string]string{
									//"": "",
								},
							},
						},
					},
				},
			},
		})
	}

	return kcs.Sort(), nil
}
