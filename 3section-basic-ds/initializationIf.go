
package main

import "fmt"

const PI float64 = 3.14159

func main() {

	// ==========================================
	// IF WITH INITIALIZATION IN GO
	// ==========================================

	// Syntax: if initialization; condition { }
	//
	// 1. Before ; -> Declare/initialize variable.
	// 2. After ;  -> Boolean condition.
	// 3. Initialization executes before condition.
	// 4. Scope -> Entire if-else statement only.
	// 5. Initialization is optional; condition must be bool.
	// 6. Supports multiple variables and shadowing.
	// 7. Remember: Initialize -> Check -> Execute.

	if x := 10; x > 5 {
		fmt.Printf("x is greater than 5: %d\n", x)
	} else {
		fmt.Printf("x is not greater than 5: %d\n", x)
	}

	// fmt.Println(x) // ERROR: x is out of scope

	// ==========================================
	// TYPE CONVERSION INSIDE IF
	// ==========================================

	// y := 4 -> int
	// PI -> float64
	// float64(y) -> Explicit int to float64 conversion.
	// Go does not implicitly convert numeric variable types.

	if y := 4; PI < float64(y) {
		fmt.Printf("PI is less than y: %f\n", PI)
	} else {
		fmt.Printf("PI is not less than y: %f\n", PI)
	}

	// ==========================================
	// MAP OPERATIONS
	// ==========================================

	newMap := map[string]int{
		"alice": 30,
		"bob":   25,
	}

	// Map lookup returns (value, exists).
	// Missing key -> zero value (0 for int), exists=false.

	age, exists := newMap["Oggy"]
	fmt.Println("Age:", age, "Exists:", exists)

	// Insert or update a key.
	newMap["Oggy"] = 35

	// = reassigns existing variables.
	age, exists = newMap["Oggy"]
	fmt.Println("Age:", age, "Exists:", exists)

	// Delete a key.
	delete(newMap, "Oggy")

	age, exists = newMap["Oggy"]
	fmt.Println("Age:", age, "Exists:", exists)

	// ==========================================
	// MAP LOOKUP WITH IF INITIALIZATION
	// ==========================================

	// Before ; -> age, exists := newMap["alice"]
	// After ;  -> exists (boolean condition)
	// age and exists are scoped to this if-else.
	// These variables shadow the outer age and exists.

	if age, exists := newMap["alice"]; exists {
		fmt.Println("Alice's age:", age)
	} else {
		fmt.Println("Alice not found")
	}

	// Outer variables remain accessible and unchanged.
	fmt.Println("Outer age:", age, "Exists:", exists)

	// ==========================================
	// MULTIPLE VARIABLES IN IF INITIALIZATION
	// ==========================================

	if a, b := 10, 20; a < b {
		fmt.Println("a is smaller than b")
	}

	// ==========================================
	// ERROR HANDLING WITH IF INITIALIZATION
	// ==========================================

	if err := saveUser(); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("User saved successfully")
	}
	
}

// Example function returning an error.
func saveUser() error {
	return nil
}
