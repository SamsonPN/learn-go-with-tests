# learn-go-with-tests

## 1. Hello World

* writing tests
    - separation of concerns: domain code vs side effects
    - Go's testing framework
        * testing package
        * `t.Helper()` to print line of where test failed
        * `t.Run()` to create subtests
        * run `go test` where your test files are
* `go mod init [project-name]` in the root directory to creates a Go module
  - this tracks module's dependencies
* declaring functions with arguments and return types
    - named return values
* if, const, and switch
* declaring variables and constants
* `TDD Process`
    - *write failing tests to see it fails*
        * helps us make sure test is relevant and is easy to understand what it's describing
    - *write smallest amount of code to make it pass to know that it's working*
    - *refactor*
        * tests ensure we safely make changes

## 2. integers

* integers, addition
* creating Testable Examples
    - test functions whose names start with Example: Example[function-name]
    - used for documenting how a function is supposed to work
    - able to view it in godoc
* godoc
    - run `godoc -http=:6060`
    - then open a web browser and go to `http://localhost:6060/pkg/` to view your documentation
