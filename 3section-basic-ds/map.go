package main

import "fmt"

func main() {
	// A map is a built-in data type in Go that associates keys with values. It is similar to a dictionary or hash table in other programming languages.
	// Maps are unordered collections of key-value pairs, where each key is unique within the map.

	// Declare and initialize a map
	myMap := make(map[string]int) // Create an empty map with string keys and int values
	myMap["Alice"] = 25
	myMap["Bob"] = 30
	myMap["Charlie"] = 35

	// Accessing values in the map
	fmt.Println("Alice's age:", myMap["Alice"])
	fmt.Println("Bob's age:", myMap["Bob"])

	// Updating a value in the map
	myMap["Alice"] = 26
	fmt.Println("Updated Alice's age:", myMap["Alice"])

	// Deleting a key-value pair from the map
	delete(myMap, "Bob")
	fmt.Println("After deleting Bob:", myMap)

	// Checking if a key exists in the map
	// here first will be the value and second will be the boolean value which will tell us if the key exists or not
	if age, exists := myMap["Charlie"]; exists {
		fmt.Println("Charlie's age:", age)
	} else {
		fmt.Println("Charlie not found in the map.")
	}
	// uper if is equivalent to below code
	// age, exists := myMap["Charlie"]

	// if exists {
	// 	fmt.Println(age)
	// }

	//The difference: In the first if approach, age and exists are scoped only to the if/else statement.

	// iterate over the map
	// first will be the key and second will be the value
	for name, age := range myMap {
		fmt.Printf("%s is %d years old.\n", name, age)
	}

	// new map
	ctMap := map[string]int{}
	fmt.Println(len(ctMap))
	ctMap["Aniket"] = 22
	fmt.Println("Age of Aniket is %d: ", ctMap["Aniket"])

	delete(ctMap, "Aniket")

	fmt.Println("Age of Aniket is %d: ", ctMap["Aniket"])

	// with initialized value

	initMap := map[string]float64{
		"PI":           3.32,
		"Avagadaro_no": 6.67,
	}
	// both are same
	// var initMap map[string]float64 = map[string]float64{
	// 	"PI":           3.32,
	// 	"Avagadaro_no": 6.67,
	// }
	fmt.Printf("map %+v", initMap)
}
