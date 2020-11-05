package comet

import "testing"

func TestRunTestSuite(t *testing.T) {
	if err := RunTestSuite("./examples/test.yml"); err != nil {
		t.Errorf("RunTestSuite() failed. error: %v", err.Error())
	}
}
