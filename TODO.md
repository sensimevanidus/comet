# TODO

* [DONE] Remove `baseURL` from the configuration. We don't need it.
* [DONE] Change the structure of configuration (move `method` and `path` under `request`)
* [DONE] Make validator names lowercase.
* [DONE] Add a new field called `variables` to configuration.
* [DONE] Implement template variables. They'll be populated via the environment variables by default.
* [DONE] `variables` overwrite environment variable values.
* [DONE] Saved values are stored in `variables` and overwrite values if they already exist.
* Implement `regex` validator.
* Add `headers` under both `request` and `response` (and remove the `type` from both `request` and `response`).
* Change the `store` to `save` and improve the `store` part so that values from the response headers can be caught as well.
    * Introduce `body` and `header`.
* Implement `response` data storage (via `save` under configuration).
* Make sure that response body is parsed with regards to the `content-type` value under the `header` section.