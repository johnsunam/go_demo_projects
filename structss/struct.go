package main

import (
	"fmt"
	"strings"
)

type person struct {
	firstname string
	lastname  string
	age       int
	fullName  string
	gender    string
}

func (per *person) getFullName() {
	per.fullName = strings.ToUpper(per.firstname + per.lastname)
}

func (per *person) getBirthYear() int {
	return 2019 - per.age
}

func (per *person) setGender(g string) {
	per.gender = g
}
func main() {
	per := person{firstname: "Ram", lastname: "kumar", age: 20}
	per.getFullName()
	fmt.Println(per.fullName)
	fmt.Println("year", per.getBirthYear())
	per.setGender("F")
	fmt.Println(per.gender)
}
