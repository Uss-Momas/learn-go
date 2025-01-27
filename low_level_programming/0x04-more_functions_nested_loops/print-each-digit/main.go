package main

import (
	"fmt"
	"math"
)

func print_number(n int) {
	count := 0
	number := n
	for number > 0 {
		number = number / 10
		count++
	}
	value := math.Pow(10, float64(count-1))
	number = n
	for number > 0 {
		fmt.Printf("%d", number/int(value))
		number = number / int(value)
	}
}

func main() {
	print_number(98)
	print_number(402)
	print_number(1024)
	print_number(0)
	print_number(-98)
}
