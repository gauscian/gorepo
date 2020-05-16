package main

import "fmt"

func main() {
	boolOne := 1 > 2
	boolZero := 0 - 2 > 1
	fmt.Printf("%t -- %t", boolOne, boolZero)
}