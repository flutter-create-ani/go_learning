
package main

import (
	"fmt"
	"os"
)

// CONCEPT 1: Basic defer
// A deferred function executes when the surrounding function is about to return.
// It does NOT execute immediately when the defer statement is encountered.
func simpleDefer() {
	fmt.Println("Function simple defer: start")

	// This print statement is scheduled to execute when simpleDefer() returns.
	defer fmt.Println("Function simple defer: deferred")

	// Normal statements execute before deferred statements.
	fmt.Println("Function simple defer: Middle")
}

// CONCEPT 2: Multiple defer statements follow LIFO order.
// LIFO = Last In, First Out (just like a stack).
// The last defer registered executes first when the function returns.
func lifoDefer() {
	fmt.Println("Function lifo defer")

	// Registered first, so this executes last among these deferred calls.
	defer fmt.Println("First Defer")

	// Registered second, so this executes before "First Defer".
	defer fmt.Println("Second Defer")

	// Normal statements execute before any deferred calls.
	fmt.Println("Function lifoSimpleDefer: Middle")
}

func main() {
	// CONCEPT 3: File handling with defer.
	// os.Create creates a file or truncates an existing file.
	// It returns a file pointer and an error.
	file, err := os.Create("./defer.txt")

	// Always check errors before using the returned resource.
	if err != nil {
		fmt.Println("Error creating file:", err)
		return // Stop execution if file creation fails.
	}

	// Schedule file.Close() to run when main() returns.
	// This ensures the file is closed even if main() returns early later.
	defer file.Close()

	// Uncomment this line to observe simpleDefer() output.
	// simpleDefer()

	// CONCEPT 4: Deferred anonymous function.
	// An anonymous function is a function without a name.
	// It is executed when main() is about to return.
	defer func() {
		fmt.Println("Before the return of main")
	}()

	// lifoDefer() executes completely before main() continues.
	// Its own deferred statements run before it returns to main().
	lifoDefer()

	// This normal statement executes before main()'s deferred functions.
	fmt.Println("Last in main")

	// When main() returns, its deferred calls execute in reverse order:
	// 1. The anonymous function prints its message.
	// 2. file.Close() closes the file.
	//    (file.Close() is deferred before the anonymous function,
	//     so it executes after the anonymous function.)
}
