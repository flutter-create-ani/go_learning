
package main

import (
	"fmt"
	"time"
)

// In Go, the 'type' keyword is used to define a new named type.

// Define a new named type based on int.
type Age int

// Define a new named type based on string.
type UserName string

// Define a new named type based on time.Time.
type Date time.Time

// In Go, a struct groups multiple fields into one data structure.
// It is similar to a class in Java, but:
// 1. Structs do not support inheritance.
// 2. Structs do not have built-in constructors.
// 3. Structs can have methods.
// 4. Structs can implement interfaces implicitly.

// Define a custom struct type User.
type User struct {
	firstName string
	lastName  string
	age       Age
	DOB       Date
}

func main() {
	// Using custom Age type
	var userAge Age = 10
	fmt.Println("User age is:", userAge)

	// Using custom UserName type
	var newUserName UserName = "Aniket Kumar"
	fmt.Println("User name is:", newUserName)

	// Creating a struct instance
	user := User{
		firstName: "Aniket",
		lastName:  "Kumar",
		age:       Age(22),
		DOB:       Date(time.Date(2003, 1, 15, 0, 0, 0, 0, time.UTC)),
	}

	// Printing the entire struct
	fmt.Println("User struct:", user)

	// Accessing individual fields
	fmt.Println("First Name:", user.firstName)
	fmt.Println("Last Name:", user.lastName)
	fmt.Println("Age:", user.age)

	// Converting custom Date type back to time.Time
	fmt.Println("DOB:", time.Time(user.DOB).Format("2006-01-02"))

	// Updating a struct field
	user.age = Age(23)
	fmt.Println("Updated age:", user.age)
}
