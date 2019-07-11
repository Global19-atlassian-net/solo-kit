package model

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

const ConversionConfigFilename = "solo-kit-conversion.json"

// SOLO-KIT Descriptors from which code can be generated

type ConversionConfig struct {
	DocsDir   string `json:"docs_dir"`
	GoPackage string `json:"go_package"`

	// set by load
	Conversions    []*Conversion
	ConversionFile string
}

type Conversion struct {
	Name     string
	Projects []*ConversionProject
}

type ConversionProject struct {
	Version         string
	NextVersion     string
	PreviousVersion string
	GoPackage       string
}

func LoadConversionConfig(path string) (ConversionConfig, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return ConversionConfig{}, err
	}

	b, err := ioutil.ReadFile(path)
	if err != nil {
		return ConversionConfig{}, err
	}
	var cc ConversionConfig
	err = json.Unmarshal(b, &cc)
	if err != nil {
		return ConversionConfig{}, err
	}
	cc.ConversionFile = path
	return cc, err
}
