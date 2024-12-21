package main

import "fmt"

func print_times_table(n int) {
	if n > 15 || n < 0 {
		fmt.Println()
	} else {
		for i := 0; i <= n; i++ {
			for j := 0; j <= n; j++ {
				switch j {
				case 0:
					fmt.Printf("%d,", i*j)
				case n:
					fmt.Printf("%4d\n", i*j)
				default:
					fmt.Printf("%4d,", i*j)
				}
			}
		}
	}
}
