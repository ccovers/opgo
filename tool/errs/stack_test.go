package errs

import (
	"errors"
	"fmt"
)

func ExampleStack() {
	stack := Stack(0)
	fmt.Println(stack)
	// Output:
}

func ExampleWithStack() {
	err := errors.New("the error")
	fmt.Println(WithStack(Trace(err)))
	// Output:
}
