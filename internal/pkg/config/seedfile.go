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
