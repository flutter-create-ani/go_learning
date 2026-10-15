package main

import "fmt"

func main() {
	tmp := 25
	if tmp > 30 {
		fmt.Println("It's hot outside.")
	} else {
		fmt.Println("It's not hot outside.")
	}
	score := 85
	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 80 {
		fmt.Println("Grade: B")
	} else if score >= 70 {
		fmt.Println("Grade: C")
	} else if score >= 60 {
		fmt.Println("Grade: D")
	} else {
		fmt.Println("Grade: F")
	}

	userAccess := map[string]bool {
		"jane": true,
		"john": false,
	}
	hasAccess, ok := userAccess["jane"]
	fmt.Println("hasAccess:", hasAccess, "ok:", ok) 
	// what is the difference between hasAccess and ok
	// hasAccess is the value associated with the key "jane" in the userAccess map, which indicates whether Jane has access (true or false).
	// ok is a boolean that indicates whether the key "jane" exists in the userAccess map. If the key exists, ok will be true; otherwise, it will be false.
	// explain the below code
	// The code checks if the key "jane" exists in the userAccess map and retrieves its value.
	// If the key exists (ok is true) and the value (hasAccess) is true, it prints "Jane has access."
	if ok && hasAccess {
		fmt.Println("Jane has access.")
	} else {
		fmt.Println("Jane does not have access.")
	}
}