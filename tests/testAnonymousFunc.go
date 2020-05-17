package main

import (
	"fmt"
)

func main() {
	noName := func() {
		fmt.Println("I do not have a name.")
	}

	giveMeName := func(s string) {
		fmt.Printf("'%s' is good. Thanks !!", s)
	}

	fmt.Println("Time to know your name")
	noName()
	fmt.Println("I will give you one as 'funk'. How is it ?")
	giveMeName("funk")
}
