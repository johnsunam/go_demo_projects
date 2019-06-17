package main

import (
	"fmt"
)

func main() {
	n := 1
	// for loop
	for i := 10; i > 0; i-- {
		n *= i
		fmt.Println("Result:", n)
	}

}
