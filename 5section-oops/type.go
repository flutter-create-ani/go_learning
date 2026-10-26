
package main

import (
	"fmt"
	"time"
)

// =====================================================
// 1. CUSTOM TYPES
// =====================================================

// The 'type' keyword creates a new named type.

// Age is a custom type based on int.
type Age int

// UserName is a custom type based on string.
type UserName string

// Date is a custom type based on time.Time.
// Note: Date is a distinct type, not an alias for time.Time.
type Date time.Time

// =====================================================
// 2. STRUCTS
// =====================================================

// A struct groups related fields into a single data structure.
// Unlike classes in Java, Go structs do not support inheritance.
// Structs can have methods and can implement interfaces implicitly.

type User struct {
	firstName string
	lastName  string
	age       Age
	DOB       Date
}

// Employee represents an employee in an organization.
//
// Fields starting with uppercase letters are exported.
// Exported fields can be accessed from other packages.
type Employee struct {
	ID        int
	FirstName string
	LastName  string
	Position  string
	Salary    string
	IsActive  bool
	JoinedAt  time.Time
}

// =====================================================
// 3. CONSTRUCTOR FUNCTION
// =====================================================

// NewEmployee creates and returns a new Employee value.
//
// Go does not have built-in constructors like Java.
// Instead, we commonly write functions such as NewEmployee.
//
// Parameters:
// id        -> employee ID
// firstName -> employee's first name
// lastName  -> employee's last name
// position  -> employee's job title
// isActive  -> whether the employee is active
//
// The function returns an Employee VALUE, not a pointer.
func NewEmployee(
	id int,
	firstName string,
	lastName string,
	position string,
	isActive bool,
) Employee {
	return Employee{
		ID:        id,
		FirstName: firstName,
		LastName:  lastName,
		Position:  position,
		IsActive:  isActive,
		JoinedAt:  time.Now(),
	}
}

// =====================================================
// 4. MAIN FUNCTION
// =====================================================

func main() {

	// -------------------------------------------------
	// A. USING CUSTOM TYPES
	// -------------------------------------------------

	// Create a variable using the custom Age type.
	var userAge Age = 10
	fmt.Println("User age is:", userAge)

	// Create a variable using the custom UserName type.
	var newUserName UserName = "Aniket Kumar"
	fmt.Println("User name is:", newUserName)

	// -------------------------------------------------
	// B. CREATING A STRUCT INSTANCE
	// -------------------------------------------------

	// Create a User struct using a struct literal.
	// A struct literal initializes the fields of a struct.
	user := User{
		firstName: "Aniket",
		lastName:  "Kumar",
		age:       Age(22),
		DOB: Date(
			time.Date(2003, 1, 15, 0, 0, 0, 0, time.UTC),
		),
	}

	// Print the entire struct.
	// The output contains the values of its fields.
	fmt.Println("User struct:", user)

	// Access individual struct fields using the dot operator.
	fmt.Println("First Name:", user.firstName)
	fmt.Println("Last Name:", user.lastName)
	fmt.Println("Age:", user.age)

	// Convert the custom Date type back to time.Time.
	// Format uses Go's reference date: Mon Jan 2 15:04:05 MST 2006.
	fmt.Println(
		"DOB:",
		time.Time(user.DOB).Format("2006-01-02"),
	)

	// Update a struct field using the dot operator.
	user.age = Age(23)
	fmt.Println("Updated age:", user.age)

	// -------------------------------------------------
	// C. CREATING AN EMPLOYEE DIRECTLY
	// -------------------------------------------------

	// Create an Employee struct directly using a struct literal.
	// All the fields below are initialized explicitly.
	jane := Employee{
		ID:        2,
		FirstName: "Aniket",
		LastName:  "Kumar",
		Position:  "Right",
		Salary:    "10000",
		IsActive:  true,
		JoinedAt:  time.Now(),
	}

	// %v prints the struct's values.
	// %+v prints the struct's values along with field names.
	fmt.Printf("Employee: %+v\n", jane)

	// Access exported fields using the dot operator.
	fmt.Println("Employee ID:", jane.ID)
	fmt.Println("First Name:", jane.FirstName)
	fmt.Println("Last Name:", jane.LastName)
	fmt.Println("Active:", jane.IsActive)

	// -------------------------------------------------
	// D. CREATING AN EMPLOYEE USING A FUNCTION
	// -------------------------------------------------

	// Call NewEmployee to create an Employee value.
	// The function initializes the fields provided to it
	// and automatically sets JoinedAt to the current time.
	joe := NewEmployee(1, "Ankit", "Kumar", "SDE 1", true)

	// Update the Salary field after creating the employee.
	joe.Salary = "100000"

	fmt.Println("Employee before pointer:", joe)

	// -------------------------------------------------
	// E. POINTERS IN GO
	// -------------------------------------------------

	// A pointer stores the memory address of a variable.
	//
	// The & operator gives us the memory address of joe.
	// joePtr now points to the original joe variable.
	//
	// Since joe is an Employee, joePtr has type *Employee.
	// The * before Employee means "pointer to Employee".
	joePtr := &joe

	// Print the pointer.
	// This prints the memory address stored in joePtr.
	fmt.Println("Employee pointer address:", joePtr)

	// Print the original Employee value.
	fmt.Println("Employee value:", joe)

	// -------------------------------------------------
	// F. DEREFERENCING A POINTER
	// -------------------------------------------------

	// The * operator can dereference a pointer.
	// *joePtr accesses the Employee value at the address
	// stored in joePtr.
	fmt.Println("Dereferenced employee:", *joePtr)

	// Go automatically dereferences struct pointers
	// when accessing fields with the dot operator.
	// Therefore, joePtr.Salary is valid syntax.
	fmt.Println("Salary through pointer:", joePtr.Salary)

	// Modify the original employee through the pointer.
	// This changes joe itself because joePtr points to joe.
	joePtr.Salary = "120000"

	// Both statements show the updated salary.
	// There is only one original joe Employee value here.
	fmt.Println("Salary through joe:", joe.Salary)
	fmt.Println("Salary through pointer:", joePtr.Salary)

	// -------------------------------------------------
	// G. POINTERS AND VALUE COPIES
	// -------------------------------------------------

	// Assigning a struct to another variable copies its value.
	// Changing the copy does not change the original.
	employeeCopy := joe
	employeeCopy.Salary = "50000"

	fmt.Println("Original salary:", joe.Salary)
	fmt.Println("Copied employee salary:", employeeCopy.Salary)

	// Assigning a pointer copies the address, not the struct.
	// Both pointers can refer to the same original Employee.
	anotherPtr := joePtr

	anotherPtr.Salary = "150000"

	// The original joe is updated because both pointers
	// refer to the same Employee value.
	fmt.Println("Salary through joe:", joe.Salary)
	fmt.Println("Salary through joePtr:", joePtr.Salary)
	fmt.Println("Salary through anotherPtr:", anotherPtr.Salary)
}
