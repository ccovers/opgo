package errs

import (
	"errors"
	"fmt"
	"strings"
)

func ExampleTrace() {
	err := Trace(errors.New("connection timeout"))
	fmt.Println("Errors:", err.Error())
	fmt.Println("Code:", err.Code())
	fmt.Println("Message:", err.Message())
	fmt.Println("Stack:", strings.HasPrefix(err.Stack(), "ccovers/opgo/tool/errs.ExampleTrace"))
	// Output:
	// Errors: connection timeout
	// Code:
	// Message:
	// Stack: true
}
