//Project Euler Problem 551
//Currently using an n^3 algorithm -- fail
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

func sumSlice(slice []int) int {
	var sum int
	for _, v := range slice {
		sum += v
	}
	return sum
}

func sumDigits(n int, base int) int {
	var sum int
	for n > 0 {
		sum += n % base
		n /= base
	}
	return sum
}

func sumSliceDigits(slice []int) int {
	var sum int
	for _, v := range slice {
		sum += sumDigits(v, 10)
	}
	return sum
}

func euler551(n int) int {
	slice := make([]int, 1, 100000)
	slice = append(slice, 1)
	for i := 1; i < 10e15; i++ {
		slice = append(slice, sumSliceDigits(slice))
	}
	return slice[len(slice)-1]
}

func main() {
	fmt.Println(sumSlice([]int{10, 2, 13, 17}))
	fmt.Println(sumDigits(1234, 10))
	fmt.Println(euler551(0))
}
