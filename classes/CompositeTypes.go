package main

import (
	"fmt"
	"strconv"
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

type Fly struct {
	name  string
	speed int
}

func structure() {
	persona := Person{Name: "Anjali Adsod", Age: 45, Gender: "Female", Color: "Semi-White"}
	fmt.Println(persona)
	fmt.Println(persona.Color)
}

func flying(f Fly) string {
	return "name of filght is " + f.name + " and it's speed is " + strconv.Itoa(f.speed) + " Kilometer / hour"
}

func main() {
	structure()
	f := Fly{"Boeing", 500}
	fmt.Println(flying(f))
}
