package main

import (
	"fmt"
)

func main() {
	slice := []int{1, 2, 3, 3, 4, 5, 6, 20, 21, 21, 45, 100}
	find := binarySearch(slice, 20, 0, len(slice)-1)
	fmt.Println("binary print of target index:", find)
}
func binarySearch(array []int, target int, lowIndex int, highIndex int) int {
	if highIndex < lowIndex {
		return -1
	}

	mid := int((highIndex + lowIndex) / 2)
	if array[mid] > target {
		return binarySearch(array, target, lowIndex, mid)
	} else if array[mid] < target {
		return binarySearch(array, target, mid+1, highIndex)
	} else {
		return mid
	}
}
