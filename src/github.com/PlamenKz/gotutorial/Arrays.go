package main

import "fmt"

//array
//slicing operations
// a[:] all elements
// a[3:] from 4th elem to end
// a[:6] first 6 items
// a[3:6] from 4th to 6th (4,5,6)

//a := []int{} -> slice

func main() {
	//3x3 array
	var threes [3][3]float32
	//go copies array if assinged
	b := threes
	//reference
	c := &threes
	fmt.Println("hello world")
	grades := [3]int{97, 85, 83}
	grades[0] = 7
	fmt.Printf("Grades %v \n", grades)
	fmt.Printf("%v", threes)
	fmt.Printf("%v", b)
	fmt.Printf("%v", c)
}
