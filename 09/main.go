package main

import "fmt"
import "math"

func isTriplet(a, b, c float64) bool {
	return math.Pow(a, 2)+math.Pow(b, 2) == math.Pow(c, 2)
}

func main() {
	for i := 1.0; i < 1000; i++ {
		for j := i + 1.0; j < 1000; j++ {
			for k := j + 1.0; k < 1000; k++ {
				if isTriplet(i, j, k) && i+j+k == 1000 {
					fmt.Println(i, j, k)
					fmt.Println(i * j * k)
					break
				}
			}
		}
	}
}
