package main

import (
	"fmt"
	"math/big"
)

func main() {
	// Initialize two big ints with the first two numbers in the sequence.
	a := big.NewInt(0)
	b := big.NewInt(1)

	var count int

	// Initialize limit as 10^999, the smallest integer with 1000 digits.
	var limit big.Int
	limit.Exp(big.NewInt(10), big.NewInt(999), nil)

	// Loop while a is smaller than 1e1000.
	for a.Cmp(&limit) < 0 {
		// Compute the next Fibonacci number, storing it in a.
		a.Add(a, b)
		count++
		// Swap a and b so that b is the next number in the sequence.
		a, b = b, a
	}
	fmt.Println(a) // 1000-digit Fibonacci number
	fmt.Println(count)
}
