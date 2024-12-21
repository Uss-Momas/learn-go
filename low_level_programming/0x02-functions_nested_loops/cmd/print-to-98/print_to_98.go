package main

import "fmt"

func print_to_98(n int) {
	if n <= 98 {
		for i := n; i <= 98; i++ {
			switch i {
			case 98:
				fmt.Printf("%d\n", i)
			default:
				fmt.Printf("%d, ", i)
			}
		}
	} else {
		for i := n; i >= 98; i-- {
			switch i {
			case 98:
				fmt.Printf("%d\n", i)
			default:
				fmt.Printf("%d, ", i)
			}
		}
	}
}
