package main

import (
	"github.com/solo-io/go-utils/log"
	"github.com/solo-io/solo-kit/pkg/code-generator/cmd"
)

//go:generate go run generate.go

func main() {

	log.Printf("starting generate")
	if err := cmd.Generate(cmd.GenerateOptions{
		// TODO joekelley don't merge this
		RelativeRoot:  "./test",
		CompileProtos: true,
		SkipGenMocks:  true,
	}); err != nil {
		log.Fatalf("generate failed!: %v", err)
	}
}
