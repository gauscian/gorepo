package main

import (
	"fmt"
)

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
	fmt.Printf("%d, %s, %t \n",x, y, z)
	z = false
	foo()
}


func foo() {
	fmt.Println("foo")
	fmt.Printf("%d, %s, %t \n",x, y, z)
}
