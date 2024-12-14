package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

func main() {
	now := time.Now().UnixNano()
	seed := rand.NewPCG(uint64(now), uint64(now))
	random := rand.New(seed)
	n := random.Int64N(now*2+1) - now

	lastDigit := n % 10

	message := ""

	if lastDigit > 5 {
		message = "and is greater than 5"
	} else if lastDigit == 0 {
		message = "and is 0"
	} else {
		message = "and is less than 6 and not 0"
	}

	fmt.Printf("Last digit of %d is %d %s\n", n, lastDigit, message)
}
