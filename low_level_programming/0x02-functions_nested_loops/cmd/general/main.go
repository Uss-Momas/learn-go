package main

import (
	"fmt"

	"example.com/general/basics"
	"example.com/general/primitives"
)

func main() {
	r := primitives.Abs(-1)
	fmt.Printf("%d\n", r)

	r = primitives.Abs(0)
	fmt.Printf("%d\n", r)

	r = primitives.Abs(1)
	fmt.Printf("%d\n", r)

	r = primitives.Abs(-98)
	fmt.Printf("%d\n", r)

	fmt.Println("========PRINT LAST DIGIT FUNCTION===========")
	basics.PrintLastDigit(98)
	basics.PrintLastDigit(0)
	r = basics.PrintLastDigit(-1024)

	fmt.Printf("%d\n", r)
}
