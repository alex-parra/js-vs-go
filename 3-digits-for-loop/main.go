package main

import "fmt"

func main() {
	digits := []int{}
	for n := 1; n <= 9; n++ {
		digits = append(digits, n)
	}

	fmt.Println("Digits:", digits)
	// Digits: [1 2 3 4 5 6 7 8 9]
}
