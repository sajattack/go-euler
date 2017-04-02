//Project Euler Problem 20
package main

import "fmt"
import "math/big"

func factorial(x *big.Int) *big.Int {
	n := big.NewInt(1)
	if x.Cmp(big.NewInt(1)) == 0 { // if n == 1
		return big.NewInt(1)
	}
	return n.Mul(x, factorial(n.Sub(x, n)))
}

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
	result := factorial(big.NewInt(100))
	fmt.Println(result)
	fmt.Println(SumBig(result, 10))
}
