package main

import (
	"fmt"
)

func main() {
	xi := []int{1, 2, 3, 4, 5}
	// Here the slice is unfurled and then 
	ret := sum(xi...)
	fmt.Println(ret)
}

func sum(slice ...int) int {
	sum := 0
	for _, v := range slice {
		sum += v
	}
	return sum
}
