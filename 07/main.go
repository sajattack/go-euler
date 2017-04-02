//Solution to Project Euler Problem 07

package main

import "fmt"
import "math"

func isPrime(n int) bool {
	if (n%2 == 0) && (n != 2) || (n%3 == 0) && (n != 3) {
		return false
	}

	for i := 1; i < int((math.Sqrt(float64(n))+1)/6.0+1); i++ {
		if n%(6*i-1) == 0 {
			return false
		} else if n%(6*i+1) == 0 {
			return false
		}
	}
	return true
}

func main() {
	var primeSlice []int
	for i := 2; len(primeSlice) < 10001; i++ {
		if isPrime(i) {
			primeSlice = append(primeSlice, i)
		}
	}
	fmt.Println(len(primeSlice))
	fmt.Println(primeSlice[len(primeSlice)-1])
}
