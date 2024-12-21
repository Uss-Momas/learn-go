package main

import "fmt"

func jack_bauer() {
	hours := 24
	minutes := 60

	for h := 0; h < hours; h++ {
		for m := 0; m < minutes; m++ {
			fmt.Printf("%02d:%02d\n", h, m)
		}
	}
}
