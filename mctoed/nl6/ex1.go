/*
Created on Sat May 16 2020
Does this
Hands on exercise
	create a func with the identifier foo that returns an int
	create a func with the identifier bar that returns an int and a string
	call both funcs
	print out their results
Primary Author - Achin Gupta
*/

package main

import (
	"fmt"
)

func main() {
	retx := foo()
	stri, inti := bar()
	fmt.Printf("foo results type = %T, value = %d \n", retx, retx)
	fmt.Printf("bar results type = %T, value = %s \n", stri, stri)
	fmt.Printf("bar results type = %T, value = %d \n", inti, inti)
}

// return one int
func foo() int {
	return 32
}

// multiple returns
func bar() (string, int) {
	return "42", 42
}
