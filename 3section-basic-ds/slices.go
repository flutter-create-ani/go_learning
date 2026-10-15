package main

import "fmt"

func main() {
	// unlike array slices are dynamic and can grow or shrink in size.
	// Slices are built on top of arrays and provide a more flexible way to work with collections of data.
	// A slice is a descriptor of an array segment. It consists of a pointer to the array, the length of the segment, and its capacity.

	// Declare a slice of integers
	var numbers []int // This is a nil slice with length 0 and capacity 0
	fmt.Println("Initial slice:", numbers)

	// Append elements to the slice
	numbers = append(numbers, 10)
	numbers = append(numbers, 20, 30)
	fmt.Println("Updated slice:", numbers)

	// [] this make array slice if [3] then this is array slice and if [] then this is slice
	names := []string{"Alice", "Bob", "Charlie"} // Declare and initialize a slice of strings
	fmt.Println("Names slice:", names)

	// Update an element in the slice
	names[1] = "David"
	names = append(names, "Eve") // Append a new element to the slice
	fmt.Println("Updated names slice:", names)

	values := append(names, "Frank", "Grace") // Append multiple elements to the slice
	fmt.Println("Values slice:", values)
	// here names only return the original slice and values return the new slice with appended values
	// concept => slices are reference types, so when you append to a slice, it may create a new underlying array if the capacity is exceeded. The original slice remains unchanged.
	fmt.Println("Names: ", names)

	items := make([]string, 3) // Create a slice of strings with length 3
	items[0] = "Item1"
	items[1] = "Item2"
	items[2] = "Item3"
	fmt.Println("Items slice:", items)

	newItems := make([]string, 2, 3) // Create a slice of strings with length 4 and capacity upto 10
	fmt.Printf("New items slice: %+v, Len : %d, Capacity: %d\n", newItems, len(newItems), cap(newItems))
	newItems = append(newItems, "Item4", "Item5") // Append elements to the slice
	fmt.Printf("New items slice after append: %+v, Len : %d, Capacity: %d\n", newItems, len(newItems), cap(newItems))
	newItems = newItems[:2] // Reslice to length 2
	fmt.Printf("New items slice after reslice: %+v, Len : %d, Capacity: %d\n", newItems, len(newItems), cap(newItems))
}