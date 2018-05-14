package main

import (
	"fmt"
)

type vehicle interface {
	start() bool
	stop() bool
	getYear() int
	getIsOn() bool
}

type car struct {
	year int
	isOn bool
}

func (c car) start() bool {
	c.isOn = true
	return c.isOn
}

func (c car) stop() bool {
	c.isOn = false
	return c.isOn
}

func (c car) getYear() int {
	return c.year
}

func (c car) getIsOn() bool {
	return c.isOn
}

func printStatus(v vehicle) {
	fmt.Printf("The car is running? %t, it was manufactured on %d", v.getIsOn(), v.getYear())
}

func main() {
	myCar := car{isOn: false, year: 2018}
	myCar.start()
	printStatus(myCar)
}
