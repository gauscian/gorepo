/*
Created on Sun May 17 2020

Use the “defer” keyword to show that a deferred func runs after the func 
containing it exits.

Primary Author - Achin Gupta
Copyright (c) 2020 Netskope
*/

package main

import "fmt"

func main() {
	fmt.Println("foo was deferred")
	defer foo()
	fmt.Println("hello there main")
}

func foo() {
	fmt.Println("foo deferred called") 
}