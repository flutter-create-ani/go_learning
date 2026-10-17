package main

import "fmt"


func modifyValue(val int) {
	val = val * 10
	fmt.Printf("Modify value %v\n", val)
}

func modifyPtr(val *int) {
	if val == nil {
		fmt.Println("val is nil")
		return
	}
	*val = *val * 10 // deferencing 
	fmt.Printf("Modify Pointer %+v\n", *val)
}
func main() {
	// age := 10
	// agePtr  := &age
	// fmt.Println("age : ", &age) // &age show the memeory address of age
	// fmt.Println("agePtr : ", &agePtr) 

	num := 10
	modifyValue(num)
	fmt.Println(num)

	modifyPtr((&num))
	fmt.Println(num)

	grade := 50
	gradePtr := &grade
	fmt.Println("Grade pointer : ", gradePtr)
	fmt.Println("Grade pointer : ", &gradePtr)
	fmt.Println("Grade pointer : ", *(gradePtr))

}