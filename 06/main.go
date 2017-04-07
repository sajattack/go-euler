package main

import "fmt"

func sumOfSquares(n int) int {
	var sum int
	for i := 1; i <= n; i++ {
		sum += i * i
	}
	return sum
}

func squareOfSum(n int) int {
	var sum int
	for i := 1; i <= n; i++ {
		sum += i
	}
	square := sum * sum
	return square
}

func main() {
	fmt.Println(squareOfSum(100) - sumOfSquares(100))
}
