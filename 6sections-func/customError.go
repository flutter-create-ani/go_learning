
package main

import (
	"errors"
	"fmt"
	"time"
)

// CONCEPT 1: Sentinel errors
// These are predefined errors that can be reused throughout the program.
// Using the same error value makes it possible to identify an error
// with errors.Is() instead of comparing error message strings.
var errorDivideByZero = errors.New("divisible by zero")
var tooLargeError = errors.New("a is too large")

// CONCEPT 2: Custom error type
// OpError stores additional information about a failure.
// Instead of returning only a string, we can include the operation,
// a numeric code, a descriptive message, and the time of failure.
type OpError struct {
	Op      string
	Code    int
	Message string
	Time    time.Time
}

// CONCEPT 3: Constructor function
// NewOrError creates and returns a pointer to an OpError.
// Returning *OpError avoids copying the entire struct and lets callers
// access the same error object.
func NewOrError(op string, code int, message string, t time.Time) *OpError {
	return &OpError{
		Op:      op,
		Code:    code,
		Message: message,
		Time:    t,
	}
}

// CONCEPT 4: Implementing the error interface
// Go's built-in error interface requires one method:
//
//     Error() string
//
// Because OpError has this method, both OpError and *OpError
// implement the error interface.
// Returning Message determines what fmt.Println(err) displays.
func (op OpError) Error() string {
	return op.Message
}

// This function demonstrates returning a custom error as the
// built-in error interface. It always returns an error in this example.
func doSomething() error {
	return NewOrError(
		"do something",
		100,
		"do something failed",
		time.Now(),
	)
}

// CONCEPT 5: Returning both a result and an error
// Go commonly returns (result, error).
// On success: return the result and nil.
// On failure: return a suitable zero value and a non-nil error.
func divide(a, b int) (int, error) {
	// Check for division by zero before performing the operation.
	if b == 0 {
		// Return the sentinel error so callers can identify it using
		// errors.Is(err, errorDivideByZero).
		return 0, errorDivideByZero
	}

	// Reject inputs where a exceeds the allowed limit.
	if a > 1000 {
		return 0, tooLargeError
	}

	// Success: return the quotient and nil to indicate no error.
	return a / b, nil
}

func main() {
	// divide returns two values:
	// val contains the result; err contains an error or nil.
	val, err := divide(10, 5)

	// Always check the error before using the result.
	if err != nil {
		// CONCEPT 6: errors.Is()
		// Check whether err matches a particular error.
		// This is safer than comparing err.Error() with a string.
		if errors.Is(err, errorDivideByZero) {
			fmt.Println("divide by zero")
		} else if errors.Is(err, tooLargeError) {
			fmt.Println("too large")
		} else {
			// Handle any error that doesn't match the known errors.
			fmt.Println(err)
		}

		// Stop main after handling the error.
		return
	}

	// If err == nil, the operation succeeded.
	fmt.Println(val)
}
