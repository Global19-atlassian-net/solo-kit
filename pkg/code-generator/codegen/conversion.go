package codegen

import (
	"bytes"
	"sort"
	"text/template"

	"github.com/solo-io/go-utils/versionutils/kubeapi"
	"github.com/solo-io/solo-kit/pkg/errors"

	"github.com/solo-io/go-utils/log"
	code_generator "github.com/solo-io/solo-kit/pkg/code-generator"
	"github.com/solo-io/solo-kit/pkg/code-generator/codegen/templates"
	"github.com/solo-io/solo-kit/pkg/code-generator/model"
)

func GenerateConversionFiles(config *model.ConversionConfig, projects []*model.Project) (code_generator.Files, error) {
	var files code_generator.Files

	sort.SliceStable(projects, func(i, j int) bool {
		vi, err := kubeapi.ParseVersion(projects[i].ProjectConfig.Version)
		if err != nil {
			return false
		}
		vj, err := kubeapi.ParseVersion(projects[j].ProjectConfig.Version)
		if err != nil {
			return false
		}
		return vi.LessThan(vj)
	})

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
			}

			var conversion *model.Conversion
			var found bool
			if conversion, found = config.Conversions[res.Name]; !found {
				conversion = &model.Conversion{
					Name:     res.Name,
					Projects: make([]*model.ConversionProject, 0, len(projects)),
				}
			}
			conversion.Projects = append(conversion.Projects, getConversionProject(index, projects))
			config.Conversions[res.Name] = conversion
		}
	}

	// TODO joekelley don't even add them
	pruneConfig(config)

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
		content, err := generateConversionFile(config, tmpl)
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

func generateConversionFile(config *model.ConversionConfig, tmpl *template.Template) (string, error) {
	buf := &bytes.Buffer{}
	if err := tmpl.Execute(buf, config); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func getConversionProject(index int, projects []*model.Project) *model.ConversionProject {
	var nextPackage, previousPackage string
	if index < len(projects)-1 {
		nextPackage = projects[index+1].ProjectConfig.Version
	}
	if index > 0 {
		previousPackage = projects[index-1].ProjectConfig.Version
	}

	return &model.ConversionProject{
		Version:         projects[index].ProjectConfig.Version,
		GoPackage:       projects[index].ProjectConfig.GoPackage,
		NextPackage:     nextPackage,
		PreviousPackage: previousPackage,
	}
}

func pruneConfig(config *model.ConversionConfig) {
	for k, v := range config.Conversions {
		if len(v.Projects) < 2 {
			delete(config.Conversions, k)
		}
	}
}
