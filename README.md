# learn-go-with-tests

## 1. Hello World


* writing tests
    - separation of concerns: domain code vs side effects
    - Go's testing framework
        * testing package
        * `t.Helper()` to print line of where test failed
        * `t.Run()` to create subtests
    - go mod init [file-name] to create a module to run `go mod`
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
