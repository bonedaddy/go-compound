package config

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

// Generate writes an empty config template to a given path
func Generate(path string) error {
	cfg := Config{}
	data, err := yaml.Marshal(&cfg)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path, data, os.ModePerm)
}

// Load is used to load the configuration object from a file
func Load(path string) (*Config, error) {
	r, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var tmpCfg Config
	if err := yaml.Unmarshal(r, &tmpCfg); err != nil {
		return nil, err
	}
	return &tmpCfg, nil
}
