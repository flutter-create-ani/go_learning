
package main

import (
	"fmt"
	"slices"
)

func main() {

	fmt.Println("----- Advance Slices -------")
	original := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Printf("Original : %v, len: %d, cap: %d\n", original, len(original), cap(original))

	s1 := original[2:5] // low : high (excluding high)

	fmt.Printf("s1 : %v, len: %d, cap: %d\n", s1, len(s1), cap(s1))
	s2 := original[:3] // low : high (high is not included) here this mean from begining to high

	fmt.Printf("s2 : %v, len: %d, cap: %d\n", s2, len(s2), cap(s2))

	s3 := original[4:] // from low to the end
	fmt.Printf("s3 : %v, len: %d, cap: %d\n", s3, len(s3), cap(s3))

	s4 := original[:]
	fmt.Printf("s4 : %v, len: %d, cap: %d\n", s4, len(s4), cap(s4))

	isExist := slices.Contains(s4, 8)

	fmt.Println("contains : ", isExist)

	s4 = slices.Insert(s4, 0, 10000)
	fmt.Printf("Modifified s4 : %v, len: %d, cap: %d\n", s4, len(s4), cap(s4))

	fmt.Printf("Original : %v, len: %d, cap: %d\n", original, len(original), cap(original))
}
