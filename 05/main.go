package main

import "fmt"

func main() {
	var i int
	for i = 1; !divisibleUpTo(i, 20); i++ {
	}
	fmt.Println(i)
}

func divisibleUpTo(n int, d int) bool {
	for i := 1; i <= d; i++ {
		if n%i != 0 {
			return false
		}
	}
	return true
}
