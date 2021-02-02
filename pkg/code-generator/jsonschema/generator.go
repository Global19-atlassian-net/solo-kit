package jsonschema

import (
	"github.com/solo-io/solo-kit/pkg/code-generator/model"
	"google.golang.org/protobuf/types/pluginpb"
)

type Generator interface {
	Convert(resource *model.Resource) ([]byte, error)
}

type GeneratorOptions struct {
	// TODO
}

type cueGenerator struct {
	// TODO
	// Cuelang implemenation of Generator
}

func (*cueGenerator) Convert(resource *model.Resource) ([]byte, error) {
	return nil, nil
}

func NewGenerator(request *pluginpb.CodeGeneratorRequest, opts *GeneratorOptions) (Generator, error) {
	return &cueGenerator{}, nil
}
