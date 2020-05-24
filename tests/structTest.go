package main

import "fmt"

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

func main() {
	p1 := person{
		first: "achin",
		last:  "gupta",
	}

	sa1 := secretAgent{
		person: person{
			last:  "gupta",
			first: "achin",
		},
		ltk: true,
	}
	sa1.person.first = "gupta"
	sa1.last = "achin"
	fmt.Println(p1)
	fmt.Println(sa1)
}
