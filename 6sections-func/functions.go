package main

import "fmt"

func greet(name string) {
	fmt.Printf("Hello, %s \n", name)
}

func add(a , b int) {
	fmt.Printf("%d + %d = %d \n", a, b, a+b)
}

func calculateArea(width float64, height float64) float64 {
	if width < 0 || height < 0 {
		fmt.Printf("Error: widht and hieght should be positive")
		return 0
	}
	return width*height
}

func fact(n int) int {
	if n <= 1 {
		return 1
	}
	return n*fact(n-1)
}

func intSeq() func() int {
	i := 0
	return func () int  {
		i++;
		return i;
	}
}

func main() {
	greet("Aniket")
	add(1 ,2)
	area := calculateArea(4, 6)
	fmt.Println(area)
	fmt.Println(fact(5))
	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	logger := func (msg string) {
		fmt.Println(msg)
	}

	logger("Hello world")
	logger("Hello world")
}