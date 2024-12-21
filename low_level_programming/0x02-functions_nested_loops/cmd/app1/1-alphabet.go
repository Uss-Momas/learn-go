package main

import "fmt"

func main() {
	print_alphabet()
	print_alphabet_x10()
}

func print_alphabet() {
	for letter := 'a'; letter <= 'z'; letter++ {
		fmt.Printf("%c", letter)
	}
	fmt.Printf("\n")
}

func print_alphabet_x10() {
	for i := 0; i < 10; i++ {
		print_alphabet()
	}
}
