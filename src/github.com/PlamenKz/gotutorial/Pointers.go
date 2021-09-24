package main

import (
	"fmt"
)

func main() {
	var a int = 42
	var b *int = &a

	fmt.Print(b)

	c := [3]int{1, 2, 3}
	d := &c[0]
	e := &c[1]
	fmt.Printf("%v,%p,%p\n", c, d, e)
}
