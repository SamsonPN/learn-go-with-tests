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

## 2. Integers

* integers, addition
* creating Testable Examples
    - test functions whose names start with Example: Example[function-name]
    - used for documenting how a function is supposed to work
    - able to view it in godoc
* godoc
    - run `godoc -http=:6060`
    - then open a web browser and go to `http://localhost:6060/pkg/` to view your documentation

## 3. Iteration

* basic `for` loop
* Benchmmark tests
  - similar to tests
  - these tests start with Benchmark[func-name]
  - testing.B gives access to b.N which determines amount of times to run the code you want to benchmark
  - to run it do `go test -bench=.` inside the same directory as the Benchmark test

## 4. Arrays and Slices

* arrays
  - an array's size is encoded in its type
  - arr := [5]int{1, 2, 3, 4, 5} is of type [5]int
    * it WILL fail the type check if you pass it into a func that accepts [4]int
* slices
  - `var slice int[]`
  - `slice := make(int[], len, capacity)`
  - slices have a *fixed capacity* but you can create new slices from old ones using `append`
    * sums = append(sums, element)
    * append returns a new slice so must be assigned to something
  - can use the python slice syntax to get slices, e.g. slice[1:] returns all values from index 1 to end of slice
* variadic functions take in any number of arguments
  - declared like so: `func foo(args ...T)`
  - args is a []T
* len() gets length of array or slice
* `go test -cover` returns % test coverage
* `reflect.DeepEqual` is useful for comparing 2 arrays/slices by their contents but reduces type-safety of code
* able to functions to variables so that they are scoped to that function

## 5. Structs, Methods, and Interfaces

* `structs`
  - a data type that allows you to bundle related data together
* `interfaces`
  - defines what methods a type has when implementing this interface
  - in Go, implementing an interface is implicit, i.e. if a type implements all the methods of this interface, then it implements that interface
* how to declare methods:
  - func (receiverName receiverType) foo() {}
  - by convention, the receiverName should be the first letter of the receiverType
  - e.g. func (r Rectangle) Area() {}
  - the receiverType can be a struct or a Named type, e.g. type MyInt int
* Table Driven Tests
  - table driven tests allow you to easily maintain and extend a suite of test cases
  - a slice of test cases defined using a struct
    * []{ name, input, expected }
  - you iterate over this slice of test cases and use t.Run using the struct.name as the test name
  - able to call individual tests using `go test -run [test-func-name]/[test-name]`
  - e.g. go test -run TestArea/Rectangle
