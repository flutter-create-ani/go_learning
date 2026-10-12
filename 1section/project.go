package main

import "fmt"

type LogLevel int

const (
	LevelTrace LogLevel = iota
	LevelDebug
	LevelInfo
	LevelWarning
	LevelError
)

var levelName = []string{"Trace", "Debug", "Info", "Warning", "Error"}


// method receiver for LogLevel type
// explains how to implement a method for a custom type in Go. 
// The receiver is specified before the method name, 
// allowing the method to access the value of the LogLevel instance it is called on. 
// This enables us to define behavior specific to the LogLevel type, 
// such as converting it to a string representation.
// here String() method is defined for LogLevel type, which returns the string representation of the log level based on its integer value.
func (l LogLevel) String() string {
	if l < LevelTrace || l > LevelError {
		return "Unknown"
	}
	return levelName[l]
}

// normal function that takes a LogLevel as an argument and prints its integer value and string representation. and doesnot return anything.
func printLogLevel(level LogLevel) {
	fmt.Printf("Log Level: %d %s\n ", level, level.String())
}

func main() {
	printLogLevel(LevelTrace)
	printLogLevel(LevelDebug)
	printLogLevel(LevelInfo)
	printLogLevel(LevelWarning)
	printLogLevel(LevelError)
}