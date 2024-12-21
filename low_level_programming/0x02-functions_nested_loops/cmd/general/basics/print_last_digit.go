package basics

import (
	"fmt"
	"example.com/general/primitives"
)

func PrintLastDigit(number int) int {
	last := number % 10
	last = primitives.Abs(last)
	fmt.Printf("%d", last)

	return last
}
