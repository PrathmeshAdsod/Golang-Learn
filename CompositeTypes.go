package main

import (
	"fmt"
)

// struct is like a class in java
type Person struct {
	Name, Gender, Color string
	Age                 int
	/*Name   string
	Age    int
	Gender string
	Color  string*/
}

func structure() {
	persona := Person{Name: "Anjali Adsod", Age: 45, Gender: "Female", Color: "Semi-White"}
	fmt.Println(persona)
	fmt.Println(persona.Color)
}

func main() {
	structure()
}
