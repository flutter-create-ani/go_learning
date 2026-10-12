
package main

import "fmt"

// Q1: What is iota in Go?
// Ans: iota is a special identifier used inside const blocks.
// It starts from 0 and increases by 1 for each new constant declaration.
//
// Q2: Why do we need iota?
// Ans: It helps us create sequential constants automatically.
// Instead of manually assigning 0, 1, 2, 3, etc.,
// Go assigns these values using iota.
//
// Q3: Does Go have enums like Java or TypeScript?
// Ans: Go does not have a native enum keyword.
// But we can create enum-like constants using iota and custom types.

// Q4: Why are we declaring "type LogLevel int"?
// Ans: We are creating a new named type called LogLevel.
// Its underlying type is int.
//
// Benefits:
// 1. Better readability: LogLevel represents log levels.
// 2. Type safety: LogLevel is distinct from ordinary int values.
// 3. We can attach methods to LogLevel.
// 4. The compiler helps prevent mixing unrelated named types.
//
// Q5: Is declaring a custom type necessary for iota?
// Ans: NO! iota works without a custom type.
// But custom types are recommended for enum-like constants.

type LogLevel int

const (
	// iota starts from 0 in each const block.

	LogLevelDebug LogLevel = iota // 0
	LogWarn                       // 1
	LogError                      // 2
	LogInfo                       // 3
	LogFatal                      // 4

	// All constants above have the LogLevel type.
	// Go automatically repeats the previous
	// constant expression and its type.
)

// Q6: Can we use iota without declaring a type?
// Ans: YES! Here, Sunday through Saturday are
// untyped integer constants.
//
// Q7: Why are we using iota + 1?
// Ans: By default, iota starts from 0.
// Adding 1 makes Sunday = 1 instead of 0.

const (
	Sunday = iota + 1 // 1
	Monday            // 2
	Tuesday           // 3
	Wednesday         // 4
	Thursday          // 5
	Friday            // 6
	Saturday          // 7
)

// Q8: Does iota reset?
// Ans: YES! Every new const block starts iota at 0.

func main() {

	fmt.Println("Days of the week:")

	fmt.Println(Sunday)    // 1
	fmt.Println(Monday)    // 2
	fmt.Println(Tuesday)   // 3
	fmt.Println(Wednesday) // 4
	fmt.Println(Thursday)  // 5
	fmt.Println(Friday)    // 6
	fmt.Println(Saturday)  // 7

	fmt.Println("Log Levels:")

	fmt.Println(LogLevelDebug) // 0
	fmt.Println(LogWarn)       // 1
	fmt.Println(LogError)      // 2
	fmt.Println(LogInfo)       // 3
	fmt.Println(LogFatal)      // 4

	// Q9: Can we assign a LogLevel constant to a LogLevel variable?
	// Ans: YES!

	var level LogLevel = LogError
	fmt.Println(level) // 2

	// Q10: Can we assign any integer to LogLevel?
	// Ans: Go allows conversion from int to LogLevel.
	// A named type does NOT restrict values to declared constants.

	level = LogLevel(100)
	fmt.Println(level) // 100

	// For strict enum validation, write a validation method.
}
