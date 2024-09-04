package main

import "fmt"

func main() {
	fmt.Println("=============LEARNING VARIABLES=============")
	var a = "Initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "apple"
	fmt.Println(f)

	// With multiple initialized values

	var i, j int = 1, 2
	var g, python, java = true, false, "no!"

	fmt.Println(i, j, g, python, java)
}