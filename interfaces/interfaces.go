package main

import (
	"fmt"
	"strconv"
)

type Flight interface {
	Speed() string
	Name() string
}

func (f Fly) Name() string {
	return "Name of the flight is " + f.name
}

func (f Fly) Speed() string {
	return "Speed is " + strconv.Itoa(f.speed)
}

type Fly struct {
	name  string
	speed int
}

func flying(fl Flight) string {
	return fl.Name() + " " + fl.Speed()
}

func main() {
	f := Fly{"Boeing", 500}
	fmt.Println(flying(f))
}
