package main

import "fmt"

func main() {
	fibonacci()
}

func fibonacci() {
	v1, v2 := 1, 2
	fmt.Printf("%d, %d, ", v1, v2)

	count := 0

	for count < 48 {
		next_value := v1 + v2
		v1, v2 = v2, next_value
		fmt.Printf("%d, ", next_value)
		count++
	}
	fmt.Println()
}
