package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		for j := i; j < 10; j++ {
			if i == 8 && j == 9 {
				fmt.Printf("%d%d", i, j)
				continue
			}
			if i != j {
				fmt.Printf("%d%d, ", i, j)
			}
		}
	}
	fmt.Println()
}
