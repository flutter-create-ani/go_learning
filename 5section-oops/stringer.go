
package main

import (
	"fmt"
)

type Employee struct {
	ID   int
	Name string
}

type BusinessPerson struct {
	ID   int
	Name string
}

type Id int

func (e Id) String() string {
	return fmt.Sprintf("ID: %d", e)
}
// Value receiver: The method receives a copy of the struct.
// Suitable for methods that only read data.
func (e Employee) GetName() string {
	return e.Name
}

func (e BusinessPerson) GetName() string {
	return e.Name
}

// fmt.Stringer is a built-in Go interface:
//
// type Stringer interface {
//     String() string
// }
//
// Any type implementing String() string automatically
// satisfies the fmt.Stringer interface.
//
// fmt.Println() automatically calls String() when the
// value implements fmt.Stringer, providing custom output.
func (e BusinessPerson) String() string {
	return fmt.Sprintf("Person[ID:%d, Name:%s]", e.ID, e.Name)
}

func main() {
	jane := BusinessPerson{
		ID:   2,
		Name: "Ankit",
	}

	var myId Id
	myId = 222
	fmt.Println(myId)
	id := Id(124)
	fmt.Println(id)
	// Automatically calls the String() method
	// because BusinessPerson implements fmt.Stringer.
	fmt.Println(jane)
}
