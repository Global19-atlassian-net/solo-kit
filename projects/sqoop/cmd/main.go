package main

import (
	"github.com/solo-io/solo-kit/pkg/utils/log"
	gatewaysetup "github.com/solo-io/solo-kit/projects/gateway/pkg/setup"
	gloosetup "github.com/solo-io/solo-kit/projects/gloo/pkg/setup"
	"github.com/solo-io/solo-kit/projects/sqoop/pkg/setup"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("err in main: %v", err.Error())
	}
}

func run() error {
	errs := make(chan error)
	go func() {
		errs <- runGloo()
	}()
	go func() {
		errs <- runGateway()
	}()
	go func() {
		errs <- runSqoop()
	}()
	return <-errs
}

func runGloo() error {
	opts, err := gloosetup.DefaultKubernetesConstructOpts()
	if err != nil {
		return err
	}
	return gloosetup.Setup(opts)
}

func runGateway() error {
	opts, err := gatewaysetup.DefaultKubernetesConstructOpts()
	if err != nil {
		return err
	}
	return gatewaysetup.Setup(opts)
}

func runSqoop() error {
	opts, err := setup.DefaultKubernetesConstructOpts()
	if err != nil {
		return err
	}
	return setup.Setup(opts)
}