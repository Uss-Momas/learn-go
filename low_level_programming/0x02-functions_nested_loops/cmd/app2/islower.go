package main

import "fmt"

func main() {
	r := isLower('H')
	fmt.Printf("%t\n", r)

	r = isLower('o')
	fmt.Printf("%t\n", r)

	r = isLower(108)
	fmt.Printf("%t\n", r)
}

func isLower(letter byte) bool {
	if letter >= 'a' && letter <= 'z' {
		return true
	}
	return false
}
