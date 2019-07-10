package templates

import (
	"text/template"
)

var ConverterTemplate = template.Must(template.New("converter").Funcs(Funcs).Parse(`package {{ .ConversionConfig.GoPackage }}

import (
	"errors"

	"github.com/solo-io/go-utils/versionutils/kubeapi"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/kube/crd"

	// TODO joekelley maybe specify path/path/pkg
	{{ range .ConversionConfig.ConvertibleResources }}
	{{ .Project.ProjectConfig.GoPackage }}
	{{ end }}
)

type UpConverter interface {
	{{ range .ConversionConfig.ConvertibleResources }}
	{{ if .NextProject }}
	From{{ upper_camel .Project.GoPackage }}To{{ upper_camel .NextProject.GoPackage }}(src *{{ .Project.GoPackage }}.{{ .Resource.Name }}) *{{ .NextProject.GoPackage }}.{{ .NextResource.Name }}
	{{ end }}
	{{ end }}
}

type DownConverter interface {
	{{ range .ConversionConfig.ConvertibleResources }}
	{{ if .PreviousProject }}
	From{{ upper_camel .Project.GoPackage }}To{{ upper_camel .PreviousProject.GoPackage }}(src *{{ .Project.GoPackage }}.{{ .Resource.Name }}) *{{ .PreviousProject.GoPackage }}.{{ .PreviousResource.Name }}
	{{ end }}
	{{ end }}
}

type {{ upper_camel .ConversionConfig.Resource.Name }}Converter interface {
	Convert(src, dst crd.SoloKitCrd) error
}

type {{ lower_camel .ConversionConfig.Resource.Name }}Converter struct {
	upConverter   UpConverter
	downConverter DownConverter
}

func New{{ upper_camel .ConversionConfig.Resource.Name }}Converter(u UpConverter, d DownConverter) crd.Converter {
	return &{{ lower_camel .ConversionConfig.Resource.Name }}Converter{
		upConverter:   u,
		downConverter: d,
	}
}

func (c *{{ lower_camel .ConversionConfig.Resource.Name }}Converter) Convert(src, dst crd.SoloKitCrd) error {
	srcVersion, err := kubeapi.ParseVersion(src.GetObjectKind().GroupVersionKind().Version)
	if err != nil {
		return err
	}
	dstVersion, err := kubeapi.ParseVersion(dst.GetObjectKind().GroupVersionKind().Version)
	if err != nil {
		return err
	}

	if srcVersion.GreaterThan(dstVersion) {
		return c.convertDown(src, dst)
	} else if srcVersion.LessThan(dstVersion) {
		return c.convertUp(src, dst)
	}
	return crd.Copy(src, dst)
}

func (c *{{ lower_camel .ConversionConfig.Resource.Name }}Converter) convertDown(src, dst crd.SoloKitCrd) error {
	if src.GetObjectKind().GroupVersionKind().Version == dst.GetObjectKind().GroupVersionKind().Version {
		return crd.Copy(src, dst)
	}

	switch t := src.(type) {
	{{ range .ConversionConfig.ConvertibleResources }}
	{{ if .PreviousProject }}
	case *{{ lower_camel .Project.GoPackage }}.{{ upper_camel .Resource.Name }}:
		return c.convertUp(c.upConverter.From{{ upper_camel .Project.GoPackage }}To{{ upper_camel .PreviousProject.GoPackage }}(t), dst)
	{{ end }}
	{{ end }}
	}
	return errors.New("unrecognized source type, this should never happen")
}

func (c *{{ lower_camel .ConversionConfig.Resource.Name }}Converter) convertUp(src, dst crd.SoloKitCrd) error {
	if src.GetObjectKind().GroupVersionKind().Version == dst.GetObjectKind().GroupVersionKind().Version {
		return crd.Copy(src, dst)
	}

	switch t := src.(type) {
	{{ range .ConversionConfig.ConvertibleResources }}
	{{ if .NextProject }}
	case *{{ lower_camel .Project.GoPackage }}.{{ upper_camel .Resource.Name }}:
		return c.convertUp(c.upConverter.From{{ upper_camel .Project.GoPackage }}To{{ upper_camel .NextProject.GoPackage }}(t), dst)
	{{ end }}
	{{ end }}
	}
	return errors.New("unrecognized source type, this should never happen")
}
`))
