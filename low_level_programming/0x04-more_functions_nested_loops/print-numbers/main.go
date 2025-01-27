package main

import "fmt"

func more_numbers() {
	for i := 0; i < 10; i++ {
		for j := 0; j <= 14; j++ {
			fmt.Printf("%d", j)
		}
		fmt.Println()
	}
}

func main() {
	more_numbers()
}
