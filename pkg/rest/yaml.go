package rest

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func readConfigurationFromFile(configFilePath string) (*testConfiguration, error) {
	content, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		return nil, err
	}

	c := &testConfiguration{}

	if err = yaml.Unmarshal(content, c); err != nil {
		return nil, err
	}

	return c, nil
}
