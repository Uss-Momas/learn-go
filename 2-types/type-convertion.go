package main

import "fmt"

func main() {
	i := -42
	j := float64(i)

	fmt.Println(j)
	
	z := uint(j)

	fmt.Println(z)
}