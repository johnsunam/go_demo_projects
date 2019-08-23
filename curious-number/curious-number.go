package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Println("Please enter the number")
	fmt.Scan(&num)
	check := checkCurious(num)
	if check {
		fmt.Println("Number is curious")
	} else {
		fmt.Println("Number isn't curious")
	}
}

func checkCurious(num int) bool {
	square := num * num
	for num > 0 {
		if num%10 != square%10 {
			return false
		}
		num /= 10
		square /= 10
	}
	return true
}
