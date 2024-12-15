package main

import "fmt"

func main() {
	for ch := 'a'; ch < 'z'; ch++ {
		fmt.Printf("%c", ch)
	}
	for ch := 65; ch <= 90; ch++ {
		fmt.Printf("%c", rune(ch))
	}
	fmt.Println()
}
