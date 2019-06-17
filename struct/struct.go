package main

import (
	"fmt"
)

type CMember struct {
	name    string
	age     int
	address string
	rank    int
}

func main() {
	cm := CMember{
		name:    "John sunam",
		age:     26,
		address: "Itahari",
		rank:    1,
	}

	cmp := &cm
	cmp.address = "Itahari, kathmandu"
	var crew []CMember
	crew = append(crew, cm)

	for i, v := range crew {
		fmt.Println(i, v.name)
	}
	fmt.Println(cm)
	var m map[string]CMember
	m = make(map[string]CMember)
	m["keving"] = cm

	m = map[string]CMember{
		"first, ":  CMember{name: "first name", address: "first address"},
		"second, ": CMember{name: "second name", address: "second address"},
	}
	for i, v := range m {
		fmt.Println("key::", i, "value::", v)
	}
}
