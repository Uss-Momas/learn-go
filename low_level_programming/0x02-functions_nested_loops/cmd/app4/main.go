package main

import "fmt"

func main() {
	r := print_sign(98)
	fmt.Printf(", %d\n", r)

	r = print_sign(0)
	fmt.Printf(", %d\n", r)

	r = print_sign(0xff)
	fmt.Printf(", %d\n", r)

	r = print_sign(-1)
	fmt.Printf(", %d\n", r)
}
