package main

import (
	"fmt"
)

func main() {
	var num, x int
	var arr []int
	fmt.Println("Enter the number:")
	fmt.Scan(&num)
	fmt.Println("Enter the part to be divided:")
	fmt.Scan(&x)
	part := num / x
	diff := num - (part * x)

	for index := 0; index < x; index++ {
		nPart := part
		if diff > 0 {
			nPart = nPart + 1
			diff = diff - 1
		}
		arr = append(arr, nPart)
	}

	fmt.Println("number is partisioned into: ", arr)
}
