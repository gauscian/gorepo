package main

import "fmt"


func init() {
	fmt.Println("initializing - just for fun")
}


func main() {
	fmt.Println("main punch of the joke")
	xi := []int{}
	xi = append(xi, 2)
	xi = append(xi, 3)
	xi = append(xi, 4)
	xi = append(xi, 5)
	

	for _, v := range xi {
		fmt.Printf("%T, %d \n", v, v)
	}
}