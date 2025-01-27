package main

import "fmt"

func print_triangle(size int) {
	if size <= 0 {
		fmt.Println()
		return
	}
	for i := 0; i < size; i++ {
		for k := size - i - 1; k > 0; k-- {
			fmt.Print(" ")
		}
		for j := 0; j < i+1; j++ {
			fmt.Print("#")
		}
		fmt.Println()
	}
}

func main() {
	print_triangle(2)
	print_triangle(10)
	print_triangle(1)
	print_triangle(0)
}
