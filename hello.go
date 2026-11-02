package main

import (
	"fmt"
	"go_tut/constants"
	"go_tut/myutils"
)

func main() {
	fmt.Println("Hello, world!")
	var ans int
	ans = myutils.Sum(1, 2, 3, 4, 5)
	fmt.Println(ans)
	fmt.Println(constants.PI)
}
