package main

import (
	"fmt"
	"time"
)

// Struct related data ko ek single type mein group karta hai.
type Employee struct {
	ID        int
	FirstName string
	LastName  string
	Position  string
	Salary    string
	IsActive  bool
	JoinedAt  time.Time
}

// Value receiver: e ko Employee ki copy milti hai.
// Isliye e ke fields change karne se original struct change nahi hota.
func (e Employee) FullName() string {
	return e.FirstName + " " + e.LastName
}

// Pointer receiver: e original Employee ka address rakhta hai.
// Isliye original struct ke fields modify ho sakte hain.
// Go automatically &jane pass kar deta hai, agar jane addressable hai.
func (e *Employee) Deactivate() {
	e.IsActive = false
}

// Pointer receiver se original employee ki joining date update hoti hai.
// Lowercase method name ka matlab hai ki method package ke bahar accessible nahi hai.
func (e *Employee) setJoinedAt(t time.Time) {
	e.JoinedAt = t
}

// Ye normal function hai, method nahi, kyunki isme receiver nahi hai.
// Pointer pass karne se original Employee modify hota hai.
func Deactivate(e *Employee) {
	e.IsActive = false
}

// Go mein built-in constructors nahi hote.
// Isliye NewEmployee jaisi function se object create karte hain.
// Ye function Employee ki value return karta hai, pointer nahi.
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

// recover() panic ko handle karta hai, lekin sirf deferred function ke andar.
// Panic recover hone ke baad calling function return kar deta hai.
func handlePanic() {
	if r := recover(); r != nil {
		fmt.Println("Recovered from panic:", r)
	}
}

func main() {
	// Struct literal se Employee create kar rahe hain.
	// JoinedAt initialize nahi kiya, isliye usme zero time value hogi.
	jane := Employee{
		ID:        2,
		FirstName: "Aniket",
		LastName:  "Kumar",
		Position:  "Right",
		Salary:    "10000",
		IsActive:  true,
	}

	fmt.Println(jane)

	// Pointer receiver hai, isliye original jane modify hoga.
	// Go internally (&jane).Deactivate() jaisa behave karta hai.
	jane.Deactivate()
	fmt.Println(jane)

	// Normal function mein pointer explicitly pass karna padta hai.
	// &jane original variable ka address deta hai.
	Deactivate(&jane)
	fmt.Println(jane)

	// Current time se 10,000,000 minutes add karke joining date update kar rahe hain.
	jane.setJoinedAt(time.Now().Add(10000000 * time.Minute))
	fmt.Println(jane.JoinedAt)

	// Constructor function se Employee create kar rahe hain.
	joeValue := NewEmployee(3, "Joe", "Smith", "Developer", true)
	fmt.Println(joeValue.FullName())

	// &Employee{} ek Employee create karke uska pointer return karta hai.
	joe := &Employee{}

	// nil ka matlab pointer kisi valid object ko point nahi kar raha.
	joe = nil
	fmt.Println(joe)

	// defer function exit ya panic ke time execute hota hai.
	// Isliye recover handler ko panic wali line se pehle register karna zaroori hai.
	defer handlePanic()

	// FullName value receiver use karta hai.
	// Nil pointer se value receiver call karne par Go ko Employee value chahiye,
	// lekin joe nil hai; isliye runtime panic hoga.
	joe.FullName()

	// Panic recover hone par bhi execution yahan se resume nahi hoga.
	fmt.Println("This line will not execute")
}