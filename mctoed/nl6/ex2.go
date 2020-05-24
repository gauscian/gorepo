/*
Created on Sun May 17 2020
create a func with the identifier foo that
	takes in a variadic parameter of type int
	pass in a value of type []int into your func (unfurl the []int)
	returns the sum of all values of type int passed in
create a func with the identifier bar that
	takes in a parameter of type []int
	returns the sum of all values of type int passed in
Primary Author - Achin Gupta
Copyright (c) 2020 Netskope
*/

package main

import (
	"fmt"
)

func main() {
	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(foo(ii...))
}

// variadic parameter
func foo(sl ...int) int {
	total := 0
	for _, v := range sl {
		total += v
	}
	return total
}
