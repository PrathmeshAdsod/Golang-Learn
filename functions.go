package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main1() {
	printName("Prathmesh")
	fmt.Println(printSistersName())
	fmt.Println(returnMultiple1())
	fmt.Println(returnMultiple2())

	fmt.Println(nakedReturn("Kamal", 20))

	fmt.Println(division(10, 0))
	fmt.Println(division(10, 5))
}

func printName(name string) {
	fmt.Println(name)
}

func printSistersName() string {
	return "Purva"
}

/* Go allows for multiple return */

// This will return first return "Hari" and not "Krishna"
// Because we are not returning only one string (string)
// This will not generate error
func returnMultiple1() string {
	return "Hari"
	return "Krishna"
}

// This can return 2 strings because we are giving (string, string)
func returnMultiple2() (string, string) {
	return "Hare", "Krishna"
}

/*
Below function wll give error
func returnMultiple1() (string, string) {
    return "Hari"
	return "Krishna"
}
*/

// NAKED Return

func nakedReturn(name string, age int) (bio string) {
	bio = "My name is " + name + " and I am " + strconv.Itoa(age) + " years old"
	return // it will return bio
}

// Erros
func division(num1, num2 int) (int, error) {
	if num2 == 0 {
		return 0, errors.New("Do not divide it by 0")
	}
	var value int = num1 / num2
	return value, nil
}
