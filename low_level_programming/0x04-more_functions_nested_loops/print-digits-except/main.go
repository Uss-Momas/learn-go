package main

import "fmt"

func print_most_numbers() {
	for c := 48; c <= 57; c++ {
		if c != '2' && c != '4' {
			fmt.Printf("%c", c)
		}
	}
	fmt.Println()
}

func main() {
	print_most_numbers()
}
