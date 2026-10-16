
package main

import "fmt"

func main() {
	// Go does not have a built-in Set data structure.
	// We use map[T]struct{} to implement a HashSet.
	// struct{} occupies zero bytes of storage.

	// Similar to HashSet<String> in Java.
	set := make(map[string]struct{})

	fmt.Println("Length of set is:", len(set))

	// Add elements to the set.
	set["aniket"] = struct{}{}
	set["ankit"] = struct{}{}
	set["aniket"] = struct{}{} // Duplicate, ignored.

	fmt.Println("Now after operation the length of set is:", len(set))

	// Check whether an element exists.
	_, exists := set["aniket"]
	fmt.Println("Does aniket exist?", exists)

	// Delete an element from the set.
	delete(set, "ankit")

	fmt.Println("After deletion:", len(set))

}
