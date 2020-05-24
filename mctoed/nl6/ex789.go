/*
* Created on Mon May 18 2020
Build and use an anonymous func 
* Primary Author - Achin Gupta
* Copyright (c) 2020 Netskope
*/
package main

import "fmt"

func main() {
	ann := func() {
		fmt.Println("I am annonymous")
	}
	runFunction(ann)
}

func runFunction(f func()) {
	fmt.Println("I was passed !! and I a variable here my address", f)
	f()
} 