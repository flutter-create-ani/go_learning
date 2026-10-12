
package main

import "fmt"

// What are values in Go?
// Values are data like numbers, strings, and booleans.
// Every value has a type.

func main() {
    // Printf prints formatted text.
    // \n is used to move to the next line.
    fmt.Printf("Hello," + " world!\n")

    // Q: Why does Println(1+1) work but Printf(1+1) fail?
    // Ans: Println accepts any value directly.
    // Printf needs a format string as its first argument.

    // Integer
    fmt.Println(1 + 1)

    // Decimal number
    fmt.Println(3.14)

    // Boolean
    fmt.Println(true, false)

    // %v prints a value in its default format.
    // %+v also shows field names for structs.
    fmt.Printf("%+v\n", []int{1, 2, 3})
}
