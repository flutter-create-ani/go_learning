package main

import "fmt"

func sum(a, b int) int {
	return a+b
}

func s(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func config(numbers ...int) {
	if len(numbers) > 0 {
		first := numbers[0]
		fmt.Println("First Numver", first)
	} else {
		fmt.Println("Default number")
	}

}

func main() {
	fmt.Println(sum(1, 2))
	fmt.Println(s(1, 2, 3, 4, 5))

	config(3,32,32,1,4)
}