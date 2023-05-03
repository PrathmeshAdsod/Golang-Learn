package main

import "fmt"

func main() {
	arr := make([]string, 0, 1)
	arr = append(arr, "Ram")
	arr = append(arr, "Sham")
	arr = append(arr, "Kam")

	fmt.Println(arr)
}
