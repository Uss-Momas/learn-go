package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

func main() {
	timestamp := time.Now().UnixNano()
	seed := rand.NewPCG(uint64(timestamp), uint64(timestamp))
	random := rand.New(seed)
	generated := random.Int64N(timestamp+timestamp+1) - timestamp

	// fmt.Printf("Random %d\n", generated)

	// verify if the number is negative or positive
	if generated < 0 {
		fmt.Printf("%d is negative\n", generated)
	} else if generated == 0 {
		fmt.Printf("%d is zero\n", generated)
	} else {
		fmt.Printf("%d is positive\n", generated)
	}
}
