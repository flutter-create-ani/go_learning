package main

import "fmt"

func main() {
	//for loop -- only way to do loops in go

	// c style loop
	for i := 1; i <= 10; i++ {
		//fmt.Println(i)
	}

	// while style loop
	k := 3
	for k > 0 {
		fmt.Println(k)
		k--
	}

	// infinite loop
	fmt.Println("-----------Infinite loop-----------")
	counter := 0
	for {
		fmt.Println(counter)
		counter++
		if counter == 10 {
			break
		}
	}

	fmt.Println("-----------Continue statement-----------")
	for i := 0; i < 8; i++ {
		if i == 2 {
			continue
		 } 
		fmt.Println(i)
	}

	fmt.Println("----------- Array -----------")
	items := [3]string{"item1", "item2", "item3"}
	for i := 0; i < len(items); i++ {
		fmt.Println(items[i])
	}

	fmt.Println("----------- Range -----------")
	for index, _ := range items {
		fmt.Println(items[index])
	}
}