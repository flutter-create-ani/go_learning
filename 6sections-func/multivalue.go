package main

import (
	"errors"
	"fmt"
	"strings"
)

func divide(a, b int) (int, error) {
	if (b == 0) {
		return 0, errors.New("Divide by zero")
	}
	return a/b, nil
}

func splitName(fullName string) (firstName, lastName string) {
	parts := strings.Split(fullName, " ")
	firstName = parts[0]
	lastName = parts[1]

	return
}

func main() {
	value , err := divide(10, 2)
	if err != nil {
		 fmt.Println(err)
	} else {
		fmt.Println(value)
	}
	first , last := splitName("Aniket kumar")
	fmt.Printf("FistName : %s lastName: %s", first, last)
}