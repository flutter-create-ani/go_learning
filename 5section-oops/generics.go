package main

import "fmt"

// Number is a Type Constraint Interface.
// | ka matlab OR hai: T int, float64 ya float32 ho sakta hai.
// Isse Type Set kehte hain.
type Number interface {
	int | float64 | float32
}

// Generic Function
// [T Number] means T ek Type Parameter hai aur Number uska constraint hai.
// T actual data type ko represent karega.
// numbers ...T means function multiple arguments accept kar sakta hai (Variadic).
// Return type T means input aur output ka type same rahega.
func sum[T Number](numbers ...T) T {

	// var totalSum := 0 this throw error why this thorwing error
	// because go is statically typed language and it needs to know the type of variable at compile time
	//
	// := se variable declare aur initialize hota hai.
	// totalSum := 0 likhne par Go iska type int infer karega.
	// Lekin generic function int, float64 aur float32 teeno accept karta hai.
	// Isliye var totalSum T use karte hain.
	// T ka actual type compile time par determine hota hai.
	var totalSum T

	for _, val := range numbers {
		totalSum += val
	}

	return totalSum
}

func main() {
	grades := []int{90, 85}
	people := []string{"Jane", "Joe", "Mark"}

	// Type Inference:
	// Compiler automatically arguments se T ka type determine karta hai.
	// sum(grades[0], grades[1]) -> T = int
	// sum(1.2, 2.3, 2.3) -> T = float64
	fmt.Println(
		len(grades),
		len(people),
		sum(grades[0], grades[1]),
		sum(1.2, 2.3, 2.3),
	)

	// here below we telling him that we are passing float64 type to the sum function
	// so here we are using type parameter to tell the compiler that we are passing float64 type to the sum function
	//
	// Explicit Type Argument:
	// Yahan hum manually T = float64 specify kar rahe hain.
	// Isse compiler ko type infer karne ki zarurat nahi hai.
	v := sum[float64](1.2, 2.3, 2.3)

	fmt.Println(v)
}
