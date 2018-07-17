package fileutils

import (
	"io/ioutil"

	"github.com/ghodss/yaml"
	"github.com/gogo/protobuf/proto"
	"github.com/pkg/errors"
	"github.com/solo-io/solo-kit/pkg/utils/protoutil"
)

func WriteToFile(filename string, pb proto.Message) error {
	jsn, err := protoutil.MarshalBytes(pb)
	data, err := yaml.JSONToYAML(jsn)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, data, 0644)
}

func ReadFileInto(filename string, v proto.Message) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return errors.Errorf("error reading file: %v", err)
	}
	jsn, err := yaml.YAMLToJSON(data)
	if err != nil {
		return err
	}
	return protoutil.UnmarshalBytes(jsn, v)
}
