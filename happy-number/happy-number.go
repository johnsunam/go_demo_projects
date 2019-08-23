package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Println("Please Enter the number:")
	fmt.Scan(&num)
	check := checkHappyNumber(num)
	if check {
		fmt.Println("It is happy number")
	} else {
		fmt.Println("It is unhappy number")
	}
}

func checkHappyNumber(num int) bool {
	var sum int
	sum = getSum(num)
	if sum == 1 {
		return true
	} else if sum == 4 {
		return false
	}
	return checkHappyNumber(sum)
}

func getSum(num int) int {
	sum := 0
	for num > 0 {
		rem := num % 10
		sum = sum + (rem * rem)
		num = num / 10
	}
	return sum
}
