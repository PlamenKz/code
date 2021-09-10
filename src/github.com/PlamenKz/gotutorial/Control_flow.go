package main

import (
	"fmt"
)

func main() {

	if true {
		fmt.Println("hi")
	}

	if true {

	} else if true {

	} else {

	}

	val := 5

	switch val {
	case 1:
		fmt.Print("1")
		fallthrough // fallthrough evaluates the case but also goes to the next case if the first one is true, doesnt break
	case 2:
		fmt.Println("2")
	default:
	}

	//basic for loop
	for i := 0; i < 5; i++ {
		fmt.Print(i)
	}

	for i := 0; i < 5; i = i + 2 {
		fmt.Print(i)
	}

	var i int
	for i < 5 {
		fmt.Print(i)
		i++
		continue
	}

	for {
		break
	}

	//breaking from mutliple loops
Loop:
	for {
		for {
			break Loop
		}
	}
	//looping through slice
	s := []int{1, 2, 3}
	//k = key v = value
	for k, v := range s {
		fmt.Println(k, v)
	}
}
