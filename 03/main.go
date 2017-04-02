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
	for i := 2; i < int(math.Sqrt(float64(n))+1); i++ {
		if isPrime(i) {
			primeSlice = append(primeSlice, i)
		}
	}
	return primeSlice
}

func main() {
	number := 600851475143
	var factors []int
	for _, prime := range primesLessThan(number) {
		if number%prime == 0 {
			factors = append(factors, prime)
		}
	}
	fmt.Println(factors[len(factors)-1])
}
