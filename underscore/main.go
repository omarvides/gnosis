package main

import (
	"fmt"
)

func main() {
	x := 20
	y := 40
	_, newY := multipleReturns(x, y)
	fmt.Println(newY)
}

func multipleReturns(x, y int) (int, int) {
	return 1 * x, 2 * y
}
