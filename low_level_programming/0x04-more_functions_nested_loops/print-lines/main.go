package main

import "fmt"

func print_line(n int) {
	if n > 0 {
		for i := 0; i < n; i++ {
			fmt.Print("_")
		}
	}
	fmt.Println()
}

func main() {
	print_line(0)
	print_line(2)
	print_line(10)
	print_line(-4)
}
