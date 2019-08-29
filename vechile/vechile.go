package main

import (
	"fmt"
	"reflect"
)

type car struct {
	numSeat int
	brand   string
}
type bike struct {
	numSeat int
	brand   string
}

type Vechicle interface {
	getSeats() int
	getBrand() string
}

func (c car) getSeats() int {
	return c.numSeat
}

func (c car) getBrand() string {
	return c.brand

}
func (b bike) getBrand() string {
	return b.brand

}

func (b bike) getSeats() int {
	return b.numSeat

}

func getValues(v Vechicle) {
	fmt.Println("type of", reflect.TypeOf(v).String())
	fmt.Println(v)
	fmt.Println(v.getSeats())
	fmt.Println(v.getBrand())
}
func main() {
	c := car{4, "nano"}
	b := bike{2, "Pulsar"}
	getValues(c)
	getValues(b)
}
