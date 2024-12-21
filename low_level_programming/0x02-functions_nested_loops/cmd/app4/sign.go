package main

import "fmt"

func print_sign(n int) int8 {
	if n > 0 {
		fmt.Print("+")
		return 1
	}
	if n == 0 {
		fmt.Print("0")
		return 0
	}
	fmt.Print("-")
	return -1
}
