/*
* Created on Mon May 18 2020
A “callback” is when we pass a func into a func as an argument. For this exercise, 
	pass a func into a func as an argument 
* Primary Author - Achin Gupta
* Copyright (c) 2020 Netskope
*/
package main

import "fmt"

func main() {
	someFunc(callback, "achin dial up repo")
}

func someFunc(callback func(string), own string) {
	fmt.Println("some Func processing-",own)
	callback(own)
}


func callback(respose string) {
	fmt.Println("Processed. ",respose)
	fmt.Println("Re encrypting dial up info")
}