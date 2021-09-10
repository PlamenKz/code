package main

import (
	"fmt"
)

func main() {
	//prints 1 in the end of the scope
	fmt.Println("")
	defer fmt.Println("1")
	fmt.Println("")
	fmt.Println("")

}
