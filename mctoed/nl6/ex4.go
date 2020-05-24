/*
* Created on Sun May 17 2020

Create a user defined struct with 
	the identifier “person”
	the fields:
		first
		last
		age
attach a method to type person with
	the identifier “speak”
	the method should have the person say their name and age
create a value of type person
call the method from the value of type person

* Primary Author - Achin Gupta
* Copyright (c) 2020 Netskope
*/

package main

import "fmt"

type person struct {
	first string
	last string
	age int
	*person
}

func (p person) speak() {
	fmt.Println("Agent name = ", p.first, p.last, p.age) 
}

func main() {
	agent := person {
		first: "achin",
		last: "gupta",
		age: 28,
	}
	agent.person = &agent
	agent.speak()
}