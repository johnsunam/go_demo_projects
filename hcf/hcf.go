package main

import (
	"fmt"
)

func main() {
	var limit, lcm int
	fmt.Println("Enter the count  of elements for  array:")
	fmt.Scan(&limit)
	var arr = make([]int, limit)
	fmt.Println("Enter the array elements")
	for i := 0; i < limit; i++ {
		var val int
		fmt.Scan(&val)
		arr[i] = val
	}
	lcm = getLcm(arr)
	fmt.Println("Here lcm for array is:", lcm)
}

func getLcm(array []int) int {
	var gcd int
	lcm := 1
	for i := 0; i < len(array); i++ {
		gcd = gcdOfNum(lcm, array[i])
		lcm = (lcm * array[i]) / gcd
	}
	return lcm
}

func gcdOfNum(val1, val2 int) int {
	var temp int
	if val1 > val2 {
		temp = val1
		val1 = val2
		val2 = temp
	}
	if val2%val1 == 0 {
		return val1
	} else {
		return gcdOfNum(val2%val1, val1)
	}

}
