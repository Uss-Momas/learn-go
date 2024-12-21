package main

import "fmt"

func times_table() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			switch j {
			case 0:
				fmt.Printf("%d,", i*j)
			case 9:
				fmt.Printf("%3d\n", i*j)
			default:
				fmt.Printf("%3d,", i*j)
			}
		}
	}
}
