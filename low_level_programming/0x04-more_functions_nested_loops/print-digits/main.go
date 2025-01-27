package main

import "fmt"

func print_numbers() {
	// WITH ASCII
	for digit := 48; digit <= 57; digit++ {
		fmt.Printf("%c", digit)
	}
	fmt.Println()
	// WITH NORMAL NUMBERS
	for i := 0; i < 10; i++ {
		fmt.Printf("%d", i)
	}
	fmt.Println()
}

func main() {
	print_numbers()
}
