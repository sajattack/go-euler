package main

import "fmt"
import "strconv"

func isPalindrome(s string) bool {
	for i, v := range s {
		if v != rune(s[len(s)-1-i]) {
			return false
		}
	}
	return true
}

func main() {
	var highest int
	for i := 999; i >= 100; i-- {
		for j := 999; j >= 100; j-- {
			if isPalindrome(strconv.Itoa(i * j)) {
				if i*j > highest {
					highest = i * j
				}
			}
		}
	}
	fmt.Println(highest)
}
