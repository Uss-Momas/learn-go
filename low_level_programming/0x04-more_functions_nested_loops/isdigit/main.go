package main

import "fmt"

func main() {
	c := '0'
	fmt.Printf("%c: %d\n", c, _isdigit(int(c)))

	c = 'a'
	fmt.Printf("%c: %d\n", c, _isdigit(int(c)))
}
