package main

import "fmt"

func main() {
	for ch := 'a'; ch <= 'z'; ch++ {
		if ch == 'e' || ch == 'q' {
			continue
		}
		fmt.Printf("%c", ch)
	}
	fmt.Println()
}
