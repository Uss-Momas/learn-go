package main

import "fmt"

func print_diagonal(n int) {
	if n > 0 {
		for i := 0; i < n; i++ {
			for j := 0; j < i; j++ {
				fmt.Print(" ")
			}
			fmt.Print("\\")
			fmt.Println()
		}
	} else {
		fmt.Println()
	}
}

func main() {
	print_diagonal(0)
	print_diagonal(2)
	print_diagonal(10)
	print_diagonal(-4)
}
