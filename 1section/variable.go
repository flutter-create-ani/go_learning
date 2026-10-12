
package main

import "fmt"

func main() {
    // Basic types: int, float64, bool, string.

    // Q: What is a zero value in Go?
    // Ans: Default value assigned when a variable is declared.
    // string = "", int = 0, bool = false, float64 = 0.0

    var greeting string
    greeting = "Hello, world!"
    fmt.Println(greeting)

    var count int
    count = 10
    fmt.Println(count)

    var isRunning bool
    isRunning = true
    fmt.Println(isRunning)

    var pi float64
    pi = 3.14
    fmt.Println(pi)

    // Declare multiple variables together.
    var firstName, lastName string
    firstName = "Aniket"
    lastName = "Kumar"
    fmt.Println(firstName, lastName)

    // Short variable declaration.
    // Go automatically detects the type from the value.
    email := "test@test.com"
    age := 30
    fmt.Println(email, age)

    // Q: What is the difference between var and :=?
    // Ans: var can declare variables with or without a value.
    // := declares and initializes a variable together.

    // Q: Can we use := outside a function?
    // Ans: No, := works only inside functions.
    // var works both inside and outside functions.

    // Go can also infer the type when using var.
    var year = 2025
    fmt.Println(year)
}
