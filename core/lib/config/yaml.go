package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func ParseYamlConfig(path string) (*ORMConfigYaml, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var config ORMConfigYaml
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
