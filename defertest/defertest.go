package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("I printed last")
	fmt.Println("I printed before")
}
