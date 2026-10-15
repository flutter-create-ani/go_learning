package main

import "fmt"

func main() {
	// in go , arrays are fixed-size collections of elements of the same type.
	//  The size of an array is determined at the time of its declaration 
	// and cannot be changed later. 
	// Arrays are useful when you know the number of elements in advance and 
	// want to store them in a contiguous block of memory.

	var numbers [2]int // Declare an array of 2 integers
	fmt.Println("Initial array: %+v\n", numbers)

	numbers[0] = 10 // Assign values to the array elements
	numbers[1] = 20
	
	fmt.Println("Updated array: %+v\n", numbers)

	primes := [5]int{2, 3, 5, 7, 11} // Declare and initialize an array of 5 integers
	fmt.Println("Primes array: %+v\n", primes)

	primes[2] = 13 // Update an element in the array
	fmt.Println("Updated primes array: %+v\n", primes)

	// Iterate over the array using a for loop
	for i := 0; i < len(primes); i++ {
		fmt.Printf("Element at index %d: %d\n", i, primes[i])
	}

	// Iterate over the array using a range loop
	for index, value := range primes {
		fmt.Printf("Element at index %d: %d\n", index, value)
	}

	var matrix [2][3]int // Declare a 2D array (matrix) of integers
	matrix[0][0] = 1
	matrix[0][1] = 2
	matrix[0][2] = 3
	matrix[1][0] = 4
	matrix[1][1] = 5
	matrix[1][2] = 6

	fmt.Println("Matrix: ", matrix)
}