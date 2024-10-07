# learn-go-with-tests

## 1. Hello World

* writing tests
    - separation of concerns: domain code vs side effects
    - Go's testing framework
        * testing package
        * `t.Helper()` to print line of where test failed
        * `t.Run()` to create subtests
        * run `go test` where your test files are
* `go mod init [project-name]` in the root directory to create a Go module
  - this tracks module's dependencies
* declaring functions with arguments and return types
    - named return values
* if, const, and switch
* declaring variables and constants
* `TDD Process`
    - *write failing tests to see it fail*
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
* able to assign functions to variables so that they are scoped to that function

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
  - it is a slice of test cases defined using a struct
    * []{ name, input, expected }
  - you iterate over this slice of test cases and use t.Run using the struct.name as the test name
  - able to call individual tests using `go test -run [test-func-name]/[test-name]`
  - e.g. go test -run TestArea/Rectangle

## 6. Pointers and Errors

* `Pointers:`
  - pointers to an object allow you to change the original object
  - functions/methods, by default, make copies of the object
  - func (w *Wallet) Deposit() {} will change the wallets balance
    * func (w Wallet) Deposit() {} will instead make a copy of Wallet and deposit it to that which does nothing
  - `Go automatically dereferences struct pointers and you do not have to use &struct to call the method`
  - if you call a method with a pointer receiver type:
    * you do not need to pass in the pointer, i.e. &Wallet.Deposit()
    * nor do you have to dereference it inside the method
      ```
      func (w *Wallet) Deposit(amount int) {
        w.balance += amount
      }
      ```
* `nil:`
  - useful to describe when a value could be missing
  - pointers can be nill
  - when a function returns a pointer, you need to make sure to check if it's nil
    * otherwise, a runtime exception will be raised
    * compiler won't help (only deals with compile time issues)
* `errors:`
  - signals a failure when calling a function/method
  - don't just check errors, handle them gracefully
  - to create a new error:
  ```
  import "errors"

  var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")
  ```
  - `go install github.com/kisielk/errcheck@latest`
	  * we ran `errcheck .` inside the directory where our tests are
	  * this essentially tells us to check for different scenarios for that error
	  * i.e. another test case to think about in this case
* t.Fatal("message")
  - stops the test when it is called
* Create new types from existing ones
  - useful when you want to apply more meaning to the values
    * i.e. we could've just used int for Wallet but this is a Wallet for bitcoin so using type Bitcoin int is more descriptive
  - can let you implement interfaces
    * e.g. we implemented the Stringer interface so we can use %s (string) when formatting Bitcoin values
    ```
    func (b Bitcoin) String() string {
	    return fmt.Sprintf("%d BTC", b)
    }
    // Output: got 10 BTC, want 20 BTC
    ```
  
## 7. Maps

* create maps
  - dictionary := map[key-type]value-type{}
  - dictionary := make(map[key-type]value-type)
  - key-type for map can only be a *comparable* type because it needs to determine if 2 keys are equal
* search for items in maps
  - map[key]
  - returns 2 values:
    1. value associated with key
    2. bool if the key actually exists in the map
      * reason being, if key does not exist in map, map[key] would return a zero-value, e.g. if map is of type map[string]int, then zero value would be 0
* update items in maps
  - just have to do map[key] = newValue
* delete items from a map
  - delete(map, key)
* more about errors:
  - how to create errors that are constants
  - writing error wrappers
  ```
    const (
      ErrNotFound         = DictionaryErr("could not find the word you were looking for")
      ErrWordExists       = DictionaryErr("word already exists in dictionary")
      ErrWordDoesNotExist = DictionaryErr("word does not exist in dictionary")
    )

    type DictionaryErr string

    // any type with an Error() string method fulfills the error interface
    func (e DictionaryErr) Error() string {
      return string(e)
    }
  ```

## 8. Dependency Injection

* we were trying to test if our Greeting function actually printed something
  - if we used fmt.Printf(), we only write to os.Stdout and that is hard to capture in our tests
  - we then looked at fmt.Printf's source code and saw that it called another function called Fprintf with os.Stdout as an argument
  - so by default fmt.Printf() uses os.Stdout but what happens if we pass in another argument instead of os.Stdout?
  - we noticed that os.Stdout implemented the io.Writer interface, so we look for another thing that implemented that to test on
* so knowing what io.Writer was, we wanted to create and `injecting our own dependency` to Fprintf instead of os.Stdout which allowed us to:
  - **Test our code**: if you can't test a function easily, it's probably because of the dependencies that are hard-wired into the function or state
    * therefore using Dependency Injection (via an interface) will help you mock out something you can control
    * in this case, the hard-wired os.Stdout as the first argument in Fprintf when calling fmt.Printf() made it hard to test if we were printing the correct thing
    * therefore, we called on Fprintf by passing in our own io.Writer which we can test separately
  - **Separate our concerns**: decoupling *where the data goes* from *how to generate it*
    * DI can help if a method/function feels like it has too many responsibilities
  - **Allow our code to be re-used in different contexts**: the first context our code can be used in is inside tests
    * but if someone else wants to test out the fucntion, they can inject their own dependencies


