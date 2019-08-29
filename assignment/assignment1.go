package main

import (
	"fmt"
)

func main() {
	var (
		word       string
		check      bool = true
		wordLength int
	)
	fmt.Println("Enter your string:")
	fmt.Scanln(&word)
	wordLength = len(word)
	for i := 0; i < wordLength; i++ {
		if word[i] != word[wordLength-(i+1)] {
			check = false
			break
		}
	}
	if check {
		fmt.Println("word is palindrom")
		rem := wordLength % 2
		div := wordLength / 2
		val := div
		if rem != 0 {
			val = div + 1
		}
		m := make(map[string]int)
		for i := div; i >= 0; i-- {
			for j := 0; j < i; j++ {
				fmt.Print(" ")
			}
			fmt.Print(word[i:div])

			fmt.Print(word[div : val+(div-i)])
			fmt.Println("")
			_, prs := m[word[i:i+1]]
			if div == i {
				if rem != 0 {
					m[word[i:i+1]] = 1
				} else {
					m[word[i:i+1]] = 2
				}
			} else {
				if prs {
					m[word[i:i+1]] = m[word[i:i+1]] + 2
				} else {
					m[word[i:i+1]] = 2
				}
			}

		}
		fmt.Println("=====", m)
	} else {
		fmt.Println("word is not palindrom")

	}

}
