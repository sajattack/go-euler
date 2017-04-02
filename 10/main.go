//Solution to Project Euler Problem 03

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

func primesLessThan(n int) []int {
	var primeSlice []int
	for i := 2; i < n; i++ {
		if isPrime(i) {
			primeSlice = append(primeSlice, i)
		}
	}
	return primeSlice
}

func sumSlice(slice []int) int {
	var sum int
	for _, v := range slice {
		sum += v
	}
	return sum
}

func main() {
	fmt.Println(sumSlice(primesLessThan(2 * 10e5)))
}
