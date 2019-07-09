package webhook

import (
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/kube/webhook/server"
)

func New(webhooks ...server.KubeWebhook) (server.Webhook, error) {
	srv := server.NewWebhook()
	for _, webhook := range webhooks {
		if err := srv.Register(webhook.Path(), webhook); err != nil {
			return nil, err
		}
	}
	return srv, nil
}
