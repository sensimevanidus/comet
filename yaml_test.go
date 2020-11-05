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

	type expectation struct {
		title string
		want  interface{}
		got   interface{}
	}

	expectations := []expectation{
		{
			title: "Test suite name",
			want:  "Service Integration Test",
			got:   config.Name,
		},
		{
			title: "Number of steps",
			want:  2,
			got:   len(config.Steps),
		},
		{
			title: "First test step's title",
			want:  "Example endpoint",
			got:   config.Steps[0].Title,
		},
		{
			title: "First test step's request method",
			want:  "GET",
			got:   config.Steps[0].Request.Method,
		},
		{
			title: "First test step's request URL",
			want:  "{{.BaseURL}}/api/v1/employees",
			got:   config.Steps[0].Request.URL,
		},
		{
			title: "First test step's response's status code",
			want:  200,
			got:   config.Steps[0].Response.Code,
		},
		{
			title: "First test step's response's content type",
			want:  "application/json",
			got:   config.Steps[0].Response.Type,
		},
	}

	for _, expectation := range expectations {
		if expectation.want != expectation.got {
			t.Errorf("%v failed. want: %+v, got: %+v", expectation.title, expectation.want, expectation.got)
		}
	}
}

func TestGetTestStepFromYAML(t *testing.T) {
	config, err := readConfigurationFromFile("./examples/test.yml")
	if err != nil {
		t.Errorf("could not read configuration file. error: %v", err.Error())
	}

	if config == nil {
		t.Errorf("test configuration could not be parsed")
	}

	step, err := getTestStepFromYAML(0, config.Name, config)
	if err != nil {
		t.Errorf("could not get test step #%d from %s", 0, config.Name)
	}

	type expectation struct {
		title string
		want  interface{}
		got   interface{}
	}

	expectations := []expectation{
		{
			title: "Test step title",
			want:  "Example endpoint",
			got:   step.Title,
		},
		{
			title: "Test step's request method",
			want:  "GET",
			got:   step.Request.Method,
		},
		{
			title: "Test step's request URL",
			want:  "http://dummy.restapiexample.com/api/v1/employees",
			got:   step.Request.URL,
		},
		{
			title: "Test step's response's status code",
			want:  200,
			got:   step.Response.Code,
		},
		{
			title: "Test step's response's content type",
			want:  "application/json",
			got:   step.Response.Type,
		},
	}

	for _, expectation := range expectations {
		if expectation.want != expectation.got {
			t.Errorf("%v failed. want: %+v, got: %+v", expectation.title, expectation.want, expectation.got)
		}
	}
}
