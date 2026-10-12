
package main

import "fmt"

// Package scope: Available in all files of this package.
const HOST string = "localhost"

func main() {
    fmt.Println(HOST)

    // Q: Why does this work if HOST is a constant?
    // Ans: := creates a new local variable named HOST.
    // It shadows the package-level constant.
    HOST := HOST + ":8080"
    fmt.Println(HOST)

    // Q: Can we declare a constant without a type?
    // Ans: Yes, Go can infer its type.

    // Q: Can we declare a constant without a value?
    // Ans: No, constants must have a value.

    // Q: Can a constant use another constant?
    // Ans: Yes.
    const AppName = "MyApp"
    const FullName = AppName + " Backend"
    fmt.Println(FullName)

    // Q: Can a constant use a variable or function call?
    // Ans: No, it must use a constant expression.

    const Pi float64 = 3.14159
    fmt.Println(Pi)

    // ==========================================
    // TYPES OF SCOPE IN GO
    // ==========================================

    // 1. PACKAGE SCOPE
    // Variables, constants, types, and functions
    // declared outside functions have package scope.
    // Example: HOST declared above main().
    // Similar to module-level declarations in JS.

    // 2. FUNCTION SCOPE (LOCAL VARIABLES)
    // Variables declared inside a function
    // are accessible only within their scope.
    // Similar to local variables in JavaScript.
    age := 20
    fmt.Println(age)

    // 3. BLOCK SCOPE
    // A block is code inside {}.
    // Variables declared inside a nested block
    // cannot be accessed outside that block.
    // Similar to JavaScript let and const.
    if true {
        city := "Delhi"
        fmt.Println(city)
    }
    // fmt.Println(city) // Error: city is undefined.

    // 4. FILE SCOPE (IMPORTS)
    // Imported package names are available
    // only in the file where they are imported.
    // Example: import "fmt" at the top.
    // Similar to JavaScript imports.

    // 5. SHADOWING
    // An inner variable can have the same name
    // as an outer variable.
    // The inner variable temporarily hides it.
    score := 100

    {
        score := 200
        fmt.Println(score) // 200
    }

    fmt.Println(score) // 100

    // 6. GLOBAL VS LOCAL VARIABLES
    // Go does not have a special global keyword.
    // Package-level variables act like globals
    // within the package.
    // Local variables are declared inside functions.

    // ==========================================
    // IMPORTANT: := VS =
    // ==========================================

    // := declares a new variable.
    // = updates an existing variable.

    number := 10

    {
        number := 20 // New variable (shadowing).
        fmt.Println(number) // 20
    }

    fmt.Println(number) // 10

    {
        number = 30 // Updates outer variable.
    }

    fmt.Println(number) // 30

    // ==========================================
    // QUICK REVISION
    // ==========================================

    // Q: What is package scope?
    // Ans: Accessible throughout the same package.

    // Q: What is function scope?
    // Ans: Local declarations belong to a function
    // or a smaller block inside it.

    // Q: What is block scope?
    // Ans: Variables declared inside {} are limited
    // to that block and its nested blocks.

    // Q: What is file scope?
    // Ans: Imported package names are file-scoped.

    // Q: What is shadowing?
    // Ans: A new inner variable hides an outer one.

    // Q: Difference between := and =?
    // Ans: := declares, = assigns.

    // Q: Can we use := outside a function?
    // Ans: No, only inside functions.

    // Q: Can we use var outside a function?
    // Ans: Yes, it creates a package-level variable.
}
