# Contributing

Contributing to the ServiceNow package will hopefully be a fairly simple and enjoyable process, but
if you have any questions, hopefully this doc will help out.

## File Structure

```$xslt
├── CONTRIBUTING.md    You are here :)
├── LICENSE            Apache License information
├── change.go          Types and funcs to interface with the change_request SNOW table
├── cmdb.go
├── service.go
├── task.go
├── time.go
└── user.go
```


## Submitting a PR

All contributions are welcome, so please feel free to submit PRs for any changes and enhancements. T
here are a minimal amount of coding standards that will help keep this project maintainable 
longer-term:

### Before Submitting
* Use GO MetaLinter for code best-practices before submitting a PR. 
  * Somewhat similar to flake, this will help keep the code-base consistent and
  easy to read.
  * Run using: `make lint` 
  * Alternatively, you can run 
    ```
    goimports -w -l *.go
    gometalinter --config=.gometalinter.json ./...
    ```
 
* Please attempt to write tests for any new functionality that is being added.  For existing functionality, always run tests
prior to building or deploying this application
  * The make functionality will generate the test output files and render HTML that can be used in the test pages.
  * Run using: `make test`
  * Review test coverage in [GitHub Pages](https://josh5276.github.io/servicenow/)
 
### Versioning
* We try to use [SymVer](https://semver.org/) for version incrementing practices by way of git tags.  



