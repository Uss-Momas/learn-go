package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i == 9 {
			fmt.Printf("%d", i)
			continue
		}
		fmt.Printf("%d,", i)
	}
	fmt.Println()
}
