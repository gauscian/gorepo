package main

import (
	"fmt"
	)

func main() {
	type x int
	// here I just created a type x with underlying type int
	var a x
	
	a = 3
	fmt.Println(a)
}
