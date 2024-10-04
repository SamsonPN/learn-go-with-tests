// no longer using the main package and have defined our own called integers
// this will group functions for working with integers like Add
package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

// these are TESTABLE EXAMPLES
// they are functions that begin with Example

// this code will cause the example to appear in the godoc documentation
// if your code changes so that the example is not valid
// the build will fail

// to see this, run: go test -v

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}

// without this Output comment, ExampleAdd would not run!
