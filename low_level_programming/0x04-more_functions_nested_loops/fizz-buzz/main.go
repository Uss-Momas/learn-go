package main

import "fmt"

func fizz_buzz() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Print("FizzBuzz")
		} else if i%5 == 0 {
			fmt.Print("Buzz")
		} else if i%3 == 0 {
			fmt.Print("Fizz")
		} else {
			fmt.Print(i)
		}

		if i == 100 {
			fmt.Println()
			break
		}
		fmt.Print(" ")
	}
}

func main() {
	fizz_buzz()
}
