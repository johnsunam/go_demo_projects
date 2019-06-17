package main

import (
	"fmt"
)

func main() {
	X := 5
	Change(&X)
	fmt.Println("The value of x: ", X)
}

func Change(X *int) {
	*X = 10
}
