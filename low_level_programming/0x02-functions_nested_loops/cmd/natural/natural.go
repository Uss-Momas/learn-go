package main

import "fmt"

func main() {
	natural()
}

func natural() {
	sum := 0

	for i := 0; i < 1024; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	fmt.Printf("%d\n", sum)
}
