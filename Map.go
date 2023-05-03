package main

import (
	"fmt"
)

func main() {
	mapp := make(map[int]int)
	for i := 1; i < 11; i++ {
		mapp[i] = i
	}
	fmt.Println(mapp)
	fmt.Println(mapp[5])
}
