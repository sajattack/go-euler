package main

import "fmt"

func main() {
	slice := make([]int, 0)
	a, b := 1, 2
	slice = append(slice, a)
	for a < 4*10e5 {
		a = a + b
		a, b = b, a
		slice = append(slice, a)
	}

	var sum int
	for _, v := range slice {
		if v%2 == 0 {
			sum += v
		}
	}
	fmt.Println(slice)
	fmt.Println(sum)
}
