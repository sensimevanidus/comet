package comet

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type testRequest struct {
	Method string                 `yaml:"method"`
	URL    string                 `yaml:"url"`
	Type   string                 `yaml:"type"`
	Body   map[string]interface{} `yaml:"body"`
}

func (r testRequest) GetBody() (io.Reader, error) {
	body, err := json.Marshal(r.Body)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(body), nil
}

type fieldValidator struct {
	Validator string ``
}

type testResponse struct {
	Code int    `yaml:"code"`
	Type string `yaml:"type"`
	Body map[string]struct {
		Validator string  `yaml:"validator,omitempty"`
		Value     *string `yaml:"value,omitempty"`
	} `yaml:"body"`
	Store []map[string]string `yaml:"store"`
}

type testStep struct {
	Title    string       `yaml:"title"`
	Request  testRequest  `yaml:"request,omitempty"`
	Response testResponse `yaml:"response,omitempty"`
}

func (s testStep) String() string {
	return s.Title
}

type testConfiguration struct {
	path      string
	Name      string            `yaml:"name"`
	Variables map[string]string `yaml:"variables"`
	Steps     []testStep        `yaml:"steps"`
}

func RunTestSuite(configFilePath string) error {
	conf, err := readConfigurationFromFile(configFilePath)
	if err != nil {
		return err
	}

	if conf == nil {
		return fmt.Errorf("make sure to provide a valid configuration file")
	}

	// storage setup (for variables)
	getStorage(conf.Name).enrichStorageWithEnvironmentVariables()

	client := &http.Client{}

	fmt.Printf("Test suite: %s\n\n", conf.Name)
	for i, _ := range conf.Steps {
		step, err := getTestStepFromYAML(i, conf.Name, conf)
		if err != nil {
			fmt.Printf("%d. [%s] %s\n", i, step.String(), err.Error())
		} else {
			ok, err := runTestStep(*step, client, conf)
			if err != nil {
				fmt.Printf("%d. [%s] %s\n", i, step.String(), err.Error())
			} else if !ok {
				fmt.Printf("%d. [%s] failed\n", i, step.String())
			} else {
				fmt.Printf("%d. [%s] ok\n", i, step.String())
			}
		}
	}

	return nil
}

func runTestStep(step testStep, client *http.Client, conf *testConfiguration) (bool, error) {
	requestBody, err := step.Request.GetBody()
	if err != nil {
		return false, err
	}

	request, err := http.NewRequest(step.Request.Method, step.Request.URL, requestBody)
	if err != nil {
		return false, err
	}

	response, err := client.Do(request)
	if err != nil {
		return false, err
	}

	if step.Response.Code != response.StatusCode {
		return false, fmt.Errorf("unexpected response code. want: %d, got: %d", step.Response.Code, response.StatusCode)
	}

	responseBody, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		return false, err
	}

	responseBodyAsJSON := make(map[string]interface{}, 0)
	if err = json.Unmarshal(responseBody, &responseBodyAsJSON); err != nil {
		return false, err
	}

	for expectedKey, expectedValueValidator := range step.Response.Body {
		value, ok := responseBodyAsJSON[expectedKey]
		if !ok {
			return false, fmt.Errorf("could not find key in response. want: %v", expectedKey)
		}

		var fieldValidator Validator

		if expectedValueValidator.Value == nil {
			fieldValidator = getValidator(expectedValueValidator.Validator)
		} else {
			fieldValidator = getValidator(expectedValueValidator.Validator, *expectedValueValidator.Value)
		}

		ok, err = fieldValidator.Validate(value)
		if err != nil {
			return false, err
		} else if !ok {
			return false, fmt.Errorf("could not validate response for key %s. want: %+v, got: %+v", expectedKey, expectedValueValidator.Value, value)
		}
	}

	// enrich variables by storing response parts, if provided
	for _, dataStoredFromResponse := range step.Response.Store {
		for storedField, storageName := range dataStoredFromResponse {
			if value, ok := responseBodyAsJSON[storedField].(string); ok {
				getStorage(conf.Name).write(storageName, value)
			}
		}
	}

	fmt.Printf("variables: %+v\n", conf.Variables)
	fmt.Printf("storage: %+v\n", *getStorage(conf.Name))

	return true, nil
}
