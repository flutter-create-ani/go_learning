
package main

import "fmt"

func demonstrateShadowing() {
    x := 10
    fmt.Println("Outer x:", x)

    {
        x := 20 // New variable shadows outer x
        fmt.Println("Inner x:", x)
    }

    fmt.Println("Outer x after block:", x)

    {
        x = 30 // Updates outer x, no shadowing
    }

    fmt.Println("Updated outer x:", x)
}
