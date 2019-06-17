package main

import (
	"fmt"
)

func main() {
	s := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("first array", s)

	s1 := s[2:4]
	fmt.Println("second array", s1)

	s2 := s[4:]
	fmt.Println("third array", s2)

	s3 := s[:5]
	fmt.Println("forth array", s3)
}
