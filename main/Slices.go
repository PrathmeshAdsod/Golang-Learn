package main

import "fmt"

func main() {
	arr := make([]string, 0, 1)
	arr = append(arr, "Ram")
	arr = append(arr, "Sham")
	arr = append(arr, "Kam")
	arr = append(arr, "Gam")
	arr = append(arr, "Pram")

	fmt.Println(arr)
	fmt.Println(arr[1:4])
	fmt.Println(arr[1:3])
	fmt.Println(arr[1:2])

}
