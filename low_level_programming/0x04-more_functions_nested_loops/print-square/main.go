package main

import "fmt"

func print_square(n int) {
	if n <= 0 {
		fmt.Println()
		return
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Print("#")
		}
		fmt.Println()
	}
}

func main() {
	print_square(2)
	print_square(10)
	print_square(0)
}
