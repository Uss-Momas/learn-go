package main

import "fmt"

func main() {
	var c int

	c = 'A'
	fmt.Printf("%c: %d\n", c, _isupper(c))
	c = 'a'
	fmt.Printf("%c: %d\n", c, _isupper(c))

	c = 'b'
	fmt.Printf("%c: %d\n", c, _isupper(c))

	c = 'Z'
	fmt.Printf("%c: %d\n", c, _isupper(c))
}
