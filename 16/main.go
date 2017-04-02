package main

import "fmt"
import "math/big"

func SumBig(n *big.Int, base int) int {
	var sum int
	i := new(big.Int).Set(n)
	b := new(big.Int).SetUint64(uint64(base))
	r := new(big.Int)
	for i.BitLen() > 0 {
		i.DivMod(i, b, r)
		sum += int(r.Uint64())
	}
	return sum
}

func main() {
	var number big.Int
	number.Exp(big.NewInt(2), big.NewInt(1000), nil)
	fmt.Println(SumBig(&number, 10))
}
