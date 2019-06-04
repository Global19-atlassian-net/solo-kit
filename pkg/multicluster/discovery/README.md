# how to create a gke service account and supply the credentials to gke discovery


export project_id=solo-public
gcloud iam --project=$project_id service-accounts create gke-cluster-discovery \
  --display-name gke-cluster-discovery
gcloud iam --project=$project_id roles create gke_cluster_discovery \
  --project $project_id \
  --title gke-cluster-discovery \
  --description "Discover GKE Clusters" \
  --permissions container.clusters.list
export service_account_email=$(gcloud iam --project=$project_id service-accounts list --filter gke-cluster-discovery --format 'value([email])')
gcloud projects add-iam-policy-binding $project_id \
  --member=serviceAccount:${service_account_email} \
  --role=projects/${project_id}/roles/gke_cluster_discovery
gcloud iam --project=$project_id service-accounts keys create \
  --iam-account $service_account_email \
  google_service_account.json

