package main

import "fmt"


func init() {
	fmt.Println("initializing jff")
}

func main() {
	array := [5]int{}
	array[0] = 1 
	array[1] = 2
	array[2] = 3
	array[3] = 4
	array[4] = 5
	// this is how you print range array
	for _, v := range array {
		fmt.Printf("type - %T data - %d \n", v, v)
	}
}