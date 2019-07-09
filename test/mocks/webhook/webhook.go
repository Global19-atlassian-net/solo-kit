package main

import (
	"context"

	"github.com/solo-io/go-utils/log"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/kube/crd"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/kube/webhook"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/kube/webhook/converter"
	"github.com/solo-io/solo-kit/test/mocks/v2alpha1"
)

type mockConverter struct {
}

func (m *mockConverter) Convert(src crd.SoloKitCrd, dst crd.SoloKitCrd) error {
	panic("implement me")
}

func main() {
	ctx := context.TODO()
	conv, err := converter.NewKubeWebhook(ctx, v2alpha1.MockResourceCrd.GroupKind(), &mockConverter{}, "crdconvert")
	if err != nil {
		log.Fatalf("%v", err)
	}
	srv, err := webhook.New(conv)
	if err != nil {
		log.Fatalf("%v", err)
	}
	err = srv.Start(ctx)
}
