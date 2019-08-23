package main

import (
	"fmt"
)

func main() {
	var firstNum, secondNum, sum int
	var arr []int
	fmt.Println("Enter your first number")
	fmt.Scan(&firstNum)
	fmt.Println("Enter your second number")
	fmt.Scan(&secondNum)
	for i := firstNum; i <= secondNum; i++ {
		check := false
		arr = getDivisor(i)
		sum = getSum(arr)
		if sum > i {
			check = checkWeirdNum(arr, i)
		}
		if check {
			fmt.Println("wired number", i)
		}
	}
}

func getDivisor(num int) []int {
	var divNum []int
	for i := 1; i < num; i++ {
		if num%i == 0 {
			divNum = append(divNum, i)
		}
	}
	return divNum
}

func checkWeirdNum(arr []int, n int) bool {
	var subset []int
	for i := 0; i < len(arr); i++ {
		var a, b []int
		var subsetSum int
		var res bool
		if i == 0 {
			subset = arr[1:]
			subsetSum = getSum(subset)
		} else {
			a = arr[:i]
			b = arr[(i + 1):]
			subset = concateSlice(a, b)
			subsetSum = getSum(subset)
		}
		if subsetSum == n {
			return false
		}
		if subsetSum > n {
			res = checkWeirdNum(subset, n)
			if !res {
				return false
			}
		}
	}
	return true
}

func getSum(arr []int) int {
	var subsetSum int
	for _, value := range arr {
		subsetSum = subsetSum + value
	}
	return subsetSum
}

func concateSlice(a []int, b []int) []int {
	var mergedArray []int
	for _, value := range a {
		mergedArray = append(mergedArray, value)
	}
	for _, value := range b {
		mergedArray = append(mergedArray, value)
	}
	return mergedArray
}
