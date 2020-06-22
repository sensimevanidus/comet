package rest

import "testing"

func TestReadConfigurationFromFile(t *testing.T) {
	config, err := readConfigurationFromFile("./examples/test.yml")
	if err != nil {
		t.Errorf("could not read configuration file. error: %v", err.Error())
	}

	if config == nil {
		t.Errorf("test configuration could not be parsed")
	}
}
