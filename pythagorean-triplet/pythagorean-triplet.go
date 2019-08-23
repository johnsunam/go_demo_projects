package main

import (
	"fmt"
)

func main() {
	var a, b, c int
	sum := 1000
	proved := false
	for a = 1; a < sum/3; a++ {
		for b = a + 1; b < sum/2; b++ {
			c = sum - a - b

			if (a*a + b*b) == c*c {
				fmt.Println("aaa", a, b, c)
				proved = true
				break
			}
		}
		if proved {
			break
		}
	}
}
