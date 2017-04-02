//Solution to Project Euler Problem 48

package main

import "fmt"
import "math/big"

func main() {
	sum := big.NewInt(0)
	for i := 1; i <= 1000; i++ {
		var power big.Int
		sum.Add(sum, power.Exp(big.NewInt(int64(i)), big.NewInt(int64(i)), nil))
	}
	fmt.Println(sum)
}
