# TODO

* [DONE] Remove `baseURL` from the configuration. We don't need it.
* [DONE] Change the structure of configuration (move `method` and `path` under `request`)
* [DONE] Make validator names lowercase.
* [DONE] Add a new field called `variables` to configuration.
* [DONE] Implement template variables. They'll be populated via the environment variables by default.
* [DONE] `variables` overwrite environment variable values.
* [DONE] Saved values are stored in `variables` and overwrite values if they already exist.
* [DONE] Update CLI output
```
# SUCCESS

## normal output
PASS
ok      github.com/sensimevanidus/comet 0.108s

## verbose output
=== RUN   TestRunTestSuite
--- PASS: TestRunTestSuite (0.00s)
=== RUN   TestStringValidatorValidate
--- PASS: TestStringValidatorValidate (0.00s)
=== RUN   TestReadConfigurationFromFile
--- PASS: TestReadConfigurationFromFile (0.00s)
=== RUN   TestGetTestStepFromYAML
--- PASS: TestGetTestStepFromYAML (0.00s)
PASS
ok      github.com/sensimevanidus/comet 0.108s

# FAILURE

## normal output
--- FAIL: TestStringValidatorValidate (0.00s)
    validator_test.go:27: string validation failed. error: given value does not match the expected one. want: random-string-valuae, got: random-string-value
    validator_test.go:31: string validation failed
FAIL
exit status 1
FAIL    github.com/sensimevanidus/comet 0.108s

## verbose output
=== RUN   TestRunTestSuite
--- PASS: TestRunTestSuite (0.00s)
=== RUN   TestStringValidatorValidate
    validator_test.go:27: string validation failed. error: given value does not match the expected one. want: random-string-valuae, got: random-string-value
    validator_test.go:31: string validation failed
--- FAIL: TestStringValidatorValidate (0.00s)
=== RUN   TestReadConfigurationFromFile
--- PASS: TestReadConfigurationFromFile (0.00s)
=== RUN   TestGetTestStepFromYAML
--- PASS: TestGetTestStepFromYAML (0.00s)
FAIL
exit status 1
FAIL    github.com/sensimevanidus/comet 0.269s
```
* Default to status code 200 when no response is provided
* Make verbose output optional
* Implement correct counters for test steps + test suite
* Implement `regex` validator.
* Add `headers` under both `request` and `response` (and remove the `type` from both `request` and `response`).
* Change the `store` to `save` and improve the `store` part so that values from the response headers can be caught as well.
    * Introduce `body` and `header`.
* Implement `response` data storage (via `save` under configuration).
* Make sure that response body is parsed with regards to the `content-type` value under the `header` section.