package config

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

type Seedfile struct {
	Name       string
	Parameters []ParamOptions
}

type ParamOptions struct {
	Variable    string
	Description string
	Optional    bool
}

func (po *ParamOptions) UnmarshalYAML(unmarshal func(interface{}) error) error {
	p := struct{
		Variable string
		Description string
		Optional *bool
	}{}
	err := unmarshal(&p)
	if err != nil {
		return err
	}

	po.Variable = p.Variable
	po.Description = p.Description
	if p.Optional == nil {
		po.Optional = false
	}

	return err
}

func ParseSeedfile(filePath string) (*Seedfile, error) {
	sf := Seedfile{}

	yamlFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		return &sf, err
	}

	err = yaml.Unmarshal(yamlFile, &sf)
	return &sf, err
}
