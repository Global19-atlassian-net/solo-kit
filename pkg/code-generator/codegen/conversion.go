package codegen

import (
	"bytes"
	"text/template"

	"github.com/solo-io/solo-kit/pkg/errors"

	"github.com/solo-io/go-utils/log"
	code_generator "github.com/solo-io/solo-kit/pkg/code-generator"
	"github.com/solo-io/solo-kit/pkg/code-generator/codegen/templates"
	"github.com/solo-io/solo-kit/pkg/code-generator/model"
)

func GenerateConversionFiles(config *model.ConversionConfig, projects []*model.Project) (code_generator.Files, error) {
	var files code_generator.Files

	for index, project := range projects {
		for _, res := range project.Resources {
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
			} else if !res.IsCustom {
				log.Printf("not generating solo-kit conversion resources for non-custom resource %v", res.Name)
				continue
			}

			var conversion *model.Conversion
			var found bool
			goType := res.CustomResource.Type
			if conversion, found = config.Conversions[goType]; !found {
				conversion = &model.Conversion{
					Name:     goType,
					Projects: make([]*model.ConversionProject, 0, len(projects)),
				}
			}
			conversion.Projects = append(conversion.Projects, getConversionProject(index, projects))
			config.Conversions[goType] = conversion
		}
	}

	fs, err := generateFilesForConversionConfig(config)
	if err != nil {
		return nil, err
	}
	files = append(files, fs...)

	return files, nil
}

func generateFilesForConversionConfig(config *model.ConversionConfig) (code_generator.Files, error) {
	var v code_generator.Files
	for name, tmpl := range map[string]*template.Template{
		"resource_converter.sk.go": templates.ConverterTemplate,
	} {
		content, err := generateResourceListFile(config, tmpl)
		if err != nil {
			return nil, errors.Wrapf(err, "internal error: processing template '%v' for resource list %v failed", tmpl.ParseName, name)
		}
		v = append(v, code_generator.File{
			Filename: name,
			Content:  content,
		})
	}
	return v, nil
}

func generateResourceListFile(config *model.ConversionConfig, tmpl *template.Template) (string, error) {
	buf := &bytes.Buffer{}
	if err := tmpl.Execute(buf, config); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func getConversionProject(index int, projects []*model.Project) *model.ConversionProject {
	var nextPackage, previousPackage string
	if index < len(projects)-2 {
		nextPackage = projects[index+1].ProjectConfig.GoPackage
	}
	if index > 0 {
		previousPackage = projects[index-1].ProjectConfig.GoPackage
	}

	return &model.ConversionProject{
		GoPackage:       projects[index].ProjectConfig.GoPackage,
		NextPackage:     nextPackage,
		PreviousPackage: previousPackage,
	}
}
