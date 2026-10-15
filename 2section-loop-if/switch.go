package main

import (
	"fmt"
	"time"
)

func main() {

	// Q1: What is a normal switch in Go?
	// Ans: A switch checks a value against different cases.
	// Go automatically stops after finding the first matching case.
	// Unlike JavaScript, we don't need to write break.

	day := "Sunday"

	fmt.Println("Today is", day)

	switch day {
	case "Monday", "Saturday":
		fmt.Println("It's Monday or Saturday.")
	case "Tuesday":
		fmt.Println("It's Tuesday.")
	case "Wednesday":
		fmt.Println("It's Wednesday.")
	case "Thursday":
		fmt.Println("It's Thursday.")
	case "Friday":
		fmt.Println("It's Friday.")
	case "Sunday":
		fmt.Println("It's Sunday, the weekend!")
	default:
		fmt.Println("Invalid day.")
	}

	// Q2: What is an expressionless switch?
	// Ans: A switch without a value checks boolean conditions.
	// It executes the first case whose condition is true.
	//
	// switch { } is the same as switch true { }.
	//
	// Here, hour is declared inside the switch.
	// We can use hour inside the switch but not outside it.

	switch hour := time.Now().Hour(); {
	case hour < 12:
		fmt.Println("Good morning!")
	case hour < 18:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")
	}

	// Q3: What is an interface in Go?
	// Ans: An interface is a type that describes methods.
	// Any type that implements those methods satisfies the interface.
	//
	// Example:
	// type Speaker interface {
	//     Speak() string
	// }
	//
	// Q4: What is interface{}?
	// Ans: interface{} is an empty interface.
	// It has no required methods, so every Go type satisfies it.
	//
	// This means interface{} can hold values like:
	// int, string, bool, float64, struct, etc.
	//
	// In modern Go, we can also write "any".
	// any is simply an alias for interface{}.
	//
	// Example:
	// var x interface{} = 10
	// x = "Hello"
	// x = true
	//
	// Q5: What is a type switch?
	// Ans: A type switch checks the actual type of a value
	// stored inside an interface.
	//
	// Example:
	// switch v := i.(type) {
	// case int:
	// case string:
	// }
	//
	// Here:
	// i       = interface variable
	// .(type) = checks the actual type of i
	// v       = variable holding the value
	//
	// IMPORTANT:
	// .(type) can only be used inside a type switch.
	//
	// Q6: Why use fmt.Printf instead of fmt.Println?
	// Ans: Printf supports format specifiers:
	// %d = integer
	// %s = string
	// %t = boolean
	// %T = type of a value
	// %v = value in default format
	// \n = new line

	checkType := func(i interface{}) {
		switch v := i.(type) {
		case int:
			fmt.Printf("Integer: %d\n", v)
		case string:
			fmt.Printf("String: %s\n", v)
		case bool:
			fmt.Printf("Boolean: %t\n", v)
		default:
			fmt.Printf("Unknown type: %T, Value: %v\n", v, v)
		}
	}

	// Q7: Can we pass different data types to the same function?
	// Ans: Yes, because checkType accepts interface{}.
	// Every Go type satisfies the empty interface.

	checkType(42)       // Integer: 42
	checkType("Hello")  // String: Hello
	checkType(true)     // Boolean: true
	checkType(3.14)     // Unknown type: float64, Value: 3.14
}