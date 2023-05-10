package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println("Jai Shree Ram", i)
	}

	// Iterate over string characters
	var str string = "HariHar"
	for i := 0; i < len(str); i++ {
		fmt.Println(string(str[i]))
	}

	// FOR LOOP LIKE WHILE LOOP
	var num = 1
	for num < 100 {
		fmt.Print(num, " ")
		num++
	}

	// FOR LOOP ON RANGE
	strs := []string{"Prathmesh", "Purva", "Anjali", "Narendra", "Ram", "Sham", "Krishna"}

	for index, number := range strs {
		fmt.Println(index, number)
	}
	for _, number := range strs {
		fmt.Println(number)
	}
}
