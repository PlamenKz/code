package main

import "fmt"

//global variables
var i int = 42
var fl float32 = 32.0

const myConst int = 50

const (
	a = iota // increment a , = 0
	b        //iota infered b = 1
	c        //c = 2

	//iota is scoped to this block
)

//grouped variables
var (
	name   string  = "name"
	number float64 = 32.0
)

var J int = 30 //variable with uppercase letter exposed outside of the package

func convertToFP(i int) float32 {
	var j float32
	j = float32(i)

	return j
}

// & = and , | = or ,^ = xor , &^ = and not

// a := 8 // 2^3
// a >> 3 // 2^3 * 2^3
// a << 3 // 2^3 / 2^3

//Arrays

func main() {
	// var flag bool = true
	// var i int
	// var i int = 70
	// i = 42
	// i = 27
	//i := 42 // cant redeclare variable in that is declared globally (i initialized)

	s := "string"
	//convert to byte array
	// b := []byte(s)
	fmt.Println(string(s[2]))
}
