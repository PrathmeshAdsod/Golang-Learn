package main

import (
	"fmt"
)

// Custom TYPES
type MY_NAME string
type AGE int

// Package Level Variable Declaration
var x, y, z int = 20, 15, 18
var a, b, c = 15, false, "Run"

// d, e, f := "Ram", 15, "Sham"  // Cannot do this here
const running string = "I am running to go far"

func main() {
	fmt.Println("x y z are", x, y, z)
	fmt.Println("a b c are", a, b, c)

	d, e, f := "Ram", 15, "Sham" // This is only allowed in functional scope
	fmt.Println("d e f are", d, e, f)

	const PI float64 = 3.14
	fmt.Println(PI)
	//running = "I am walking now"  // Will raise error because it is const
	fmt.Println(running)

	typeConversion()
	Pointers()
}

func typeConversion() {
	var a int = 5
	var b float64 = 8.855
	fmt.Println(float64(a))
	fmt.Println(int(b))

	var c MY_NAME = "Anjali"
	var d int = int(c[0])

	fmt.Println(d)
}

func Pointers() {
	var a int = 10

	// referencing using &
	var aa *int = &a // aa points to a's memory address

	// Defrencing using *
	var aaa int = *aa
	fmt.Println(aaa, "'s memory address is ", (aa))

}
