package comet

import (
	"bytes"
	"io/ioutil"
	"text/template"

	"gopkg.in/yaml.v2"
)

func readConfigurationFromFile(configFilePath string) (*testConfiguration, error) {
	content, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		return nil, err
	}

	c := &testConfiguration{
		path: configFilePath,
	}

	if err = yaml.Unmarshal(content, c); err != nil {
		return nil, err
	}

	initStorage(c)

	return c, nil
}

func getTestStepFromYAML(stepIndex int, testName string, config *testConfiguration) (*testStep, error) {
	tmpl, err := template.ParseFiles(config.path)
	if err != nil {
		return nil, err
	}

	s := getStorage(testName)

	var tmplBytes bytes.Buffer
	tmpl.Execute(&tmplBytes, s.data)

	c2 := &testConfiguration{}
	if err = yaml.Unmarshal(tmplBytes.Bytes(), c2); err != nil {
		return nil, err
	}

	return &c2.Steps[stepIndex], nil
}
