package main

import (
	"fmt"
)

func main() {
	var first [10]int
	fmt.Println("First array:", first)
	first[4] = 20
	fmt.Println("First array:", first)

	second := [5]int{0, 24, 9, 8, 10}
	fmt.Println(second)

	var matrix [2][2]int
	fmt.Println("matrix:", matrix)
}
