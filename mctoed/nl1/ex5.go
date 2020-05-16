package main

import (
	"fmt"
)

func main() {
	type hotdog int
	var x hotdog
	x = 42
	var y int
	fmt.Printl(x)
	y = int(x)
	fmt.Printf("Hi the value for y is = %d\n", y)
}
