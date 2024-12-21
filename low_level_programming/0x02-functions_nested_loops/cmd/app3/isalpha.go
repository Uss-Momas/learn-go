package main

import "fmt"

func main() {
	r := _isalpha('H')
	_putchar(r)
	r = _isalpha('o')
	_putchar(r)
	r = _isalpha(108)
	_putchar(r)
	r = _isalpha(';')
	_putchar(r)

	fmt.Println()
}

func _putchar(value bool) {
	fmt.Printf("%t", value)
}

func _isalpha(letter byte) bool {
	if letter >= 'A' && letter <= 'Z' {
		return true
	}
	return false
}
