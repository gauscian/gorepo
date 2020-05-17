package main

import (
	"fmt"
)

type animal interface {
	eat()
}

type snake struct {
	name string
}

func (s snake) eat() {
	fmt.Printf("I am a snake, my name is %s and I swallow \n", s.name)
}

type dog struct {
	name string
}

func (d dog) eat() {
	fmt.Printf("I am a dog, my name is %s and I chew\n", d.name)
}

func howtoeat(a animal) {
	a.eat()
}

func main() {
	s := snake{
		name: "slimy",
	}

	d := dog{
		name: "bruno",
	}

	howtoeat(s)
	howtoeat(d)
}
