package main

import (
	"fmt"
)

func main() {
	var arr = [5]int{1, 2, 1, 2, 1}
	fmt.Println(arr)

	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	// Arrays is pass by value in Go
	// So the modification in another function will
	// not change array content if we don not use pointers

	/*
		update(arr, 2, 55)

		for i := 0; i < len(arr); i++ {
			fmt.Print(arr[i], " ")
		}

		fmt.Println()
	*/

	// When use P O I N T E R S
	update(&arr, 2, 10)

	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i], " ")
	}
}

func update(arr *[5]int, index int, value int) {
	(*arr)[index] = value
}
