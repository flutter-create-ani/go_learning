package main

import (
	"fmt"
)

// Go mein implements keyword nahi hota.
// Jo type interface ke saare methods implement karta hai,
// woh automatically interface satisfy karta hai.

type Person interface {
	// Interface define karta hai ki type kya kar sakta hai.
	GetName() string
}

type Employee struct {
	ID   int
	Name string
}

type BusinessPerson struct {
	ID   int
	Name string
}

// Value receiver: method ko struct ki copy milti hai.
// Sirf data read karne ke liye suitable hai.
func (e Employee) GetName() string {
	return e.Name
}

func (e BusinessPerson) GetName() string {
	return e.Name
}

// Interface parameter use karne se function multiple types
// ke saath kaam kar sakta hai (polymorphism).
func displayPerson(p Person) {
	fmt.Println(p.GetName())
}

// Compile-time checks: ensure both types implement Person.
var _ Person = Employee{}
var _ Person = BusinessPerson{}
// The _ means: "I don't need to store or use this variable,
//  but I want the compiler to check that Employee{} implements the Person interface."
// 1. Why do we use _?
// Normally, you declare a variable like this:
// var p Person = Employee{}
// Here, p stores the value as a Person interface, and you can use it:
// fmt.Println(p.GetName())
// But if you only want the compiler to check interface implementation, you don't need a variable. So you write:
// var _ Person = Employee{}
// The compiler checks the assignment, but the value is discarded.

func main() {
	joe := Employee{
		ID:   1,
		Name: "Aniket",
	}

	jane := BusinessPerson{
		ID:   2,
		Name: "Ankit",
	}

	// Dono types same function use kar sakte hain.
	displayPerson(joe)
	displayPerson(jane)
}