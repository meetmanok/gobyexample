package main

/*
In Go it’s idiomatic to communicate errors via an explicit, separate return value. This contrasts with the exceptions used in languages like Java and Ruby and the overloaded single result / error value sometimes used in C. Go’s approach makes it easy to see which functions return errors and to handle them using the same language constructs employed for any other, non-error tasks.

See this great post (http://blog.golang.org/2011/07/error-handling-and-go.html) on the Go blog for more on error handling.
*/

import (
	"errors"
	"fmt"
)

/*
By convention, errors are the last return value and have type error, a built-in interface.

errors.New constructs a basic error value with the given error message.

A nil value in the error position indicates that there was no error.
 */
func f1(arg int) (int, error) {
	if arg == 5 {
		return -1, errors.New("Got invalid number = '5'")
	}
	return arg, nil
}


/*
It’s possible to use custom types as errors by implementing the Error() method on them. Here’s a variant on the example above that uses a custom type to explicitly represent an argument error.
 */
type argError struct {
	arg int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

/*
In this case we use &argError syntax to build a new struct, supplying values for the two fields arg and prob.
 */
func f2(arg int) (int, error) {
	if arg == 5 {
		return -1, &argError{arg, "Got invalid number = '5'"}
	}
	return arg, nil
}

/*
The two loops below test out each of our error-returning functions. Note that the use of an inline error check on the if line is a common idiom in Go code.
 */
func main() {
	for _, v := range []int{3,4,5,6,7} {
		if arg, err := f1(v); err != nil {
			fmt.Println("f1 failed:", err, err.Error())
		} else {
			fmt.Println("f1 worked:", arg)
		}
	}

	fmt.Println("")

	for _, v := range []int{3,4,5,6,7} {
		if arg, err := f2(v); err != nil {
			fmt.Println("f2 failed:", err, err.Error())
		} else {
			fmt.Println("f2 worked:", arg)
		}
	}
}