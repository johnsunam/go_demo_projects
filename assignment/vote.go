package main

import (
	"fmt"
)

func main() {
	var (
		name string
		age  int
	)
	fmt.Println("Enter name:")
	fmt.Scanln(&name)
	fmt.Println("Enter age:")
	fmt.Scanln(&age)
	if age > 18 {
		fmt.Println("You are allowed to vote")
		for index := 13; index < 42; index++ {
			num := index % 2
			if num == 0 {
				fmt.Println(index)
			}
		}
	} else {
		fmt.Println("You are not allowed to vote")

		for index := 13; index < 42; index++ {
			num := index % 2
			if num != 0 {
				fmt.Println(index)
			}
		}
	}
}
