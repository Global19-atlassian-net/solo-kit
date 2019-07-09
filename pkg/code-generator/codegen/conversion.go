package codegen

import (
	"bytes"
	"text/template"

	"github.com/solo-io/solo-kit/pkg/errors"

	"github.com/iancoleman/strcase"
	"github.com/solo-io/go-utils/log"
	code_generator "github.com/solo-io/solo-kit/pkg/code-generator"
	"github.com/solo-io/solo-kit/pkg/code-generator/codegen/templates"
	"github.com/solo-io/solo-kit/pkg/code-generator/model"
)

func GenerateConversionFiles(projects []*model.Project) (code_generator.Files, error) {
	var files code_generator.Files
	// GroupKind -> Resource
	convertibleResources := make(map[string][]*model.ConvertibleResource)

	for _, project := range projects {
		for _, res := range project.Resources {
			// TODO break this out for reuse
			// only generate files for the resources in our group, otherwise we import
			if !project.ProjectConfig.IsOurProto(res.Filename) && !res.IsCustom {
				log.Printf("not generating solo-kit "+
					"clients for resource %v.%v, "+
					"resource proto package must match project proto package %v", res.ProtoPackage, res.Name, project.ProtoPackage)
				continue
			} else if res.IsCustom && res.CustomResource.Imported {
				log.Printf("not generating solo-kit "+
					"clients for resource %v.%v, "+
					"custom resources from a different project are not generated", res.GoPackage, res.Name, project.ProjectConfig.GoPackage)
				continue
			}

			if res.IsCustom {
				//var group string
				//if res.Project.ProjectConfig.CrdGroupOverride != "" {
				//	group = res.Project.ProjectConfig.CrdGroupOverride
				//} else {
				//	group = res.Project.ProtoPackage
				//}
				kind := res.CustomResource.Type
				//gk := group+kind
				convertibleResources[kind] = append(convertibleResources[kind], res)
			}
		}
	}

	for _, resources := range convertibleResources {
		fs, err := generateFilesForResourceList(resources)
		if err != nil {
			return nil, err
		}
		files = append(files, fs...)
	}

	return files, nil
}

func generateFilesForResourceList(resources []*model.ConvertibleResource) (code_generator.Files, error) {
	var v code_generator.Files
	for suffix, tmpl := range map[string]*template.Template{
		"_converter.sk.go": templates.ConverterTemplate,
	} {
		content, err := generateResourceListFile(resources, tmpl)
		if err != nil {
			return nil, errors.Wrapf(err, "internal error: processing template '%v' for resource list %v failed", tmpl.ParseName, resources[0].Resource.Name)
		}
		v = append(v, code_generator.File{
			Filename: strcase.ToSnake(resources[0].Resource.Name) + suffix,
			Content:  content,
		})
	}
	return v, nil
}

func generateResourceListFile(resources []*model.ConvertibleResource, tmpl *template.Template) (string, error) {
	buf := &bytes.Buffer{}
	if err := tmpl.Execute(buf, resources); err != nil {
		return "", err
	}
	return buf.String(), nil
}
