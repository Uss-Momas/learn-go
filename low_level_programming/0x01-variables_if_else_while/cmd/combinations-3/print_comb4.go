package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		for j := i; j < 10; j++ {
			for k := j; k < 10; k++ {
				if i == 7 && j == 8 && k == 9 {
					fmt.Printf("%d%d%d", i, j, k)
					break
				}
				if i != j && i != k && j != k {
					fmt.Printf("%d%d%d, ", i, j, k)
				}
			}
		}
	}
	fmt.Println()
}
