package main

import "fmt"

func main() {
	xi := []int{}
	// appending some values
	for i := 42; i <= 51; i++ {
		xi = append(xi, i)
	}

	// append 52
	xi = append(xi, 52)

	// print
	fmt.Println(xi)

	// append
	xi = append(xi, 53, 54, 55)

	// print
	fmt.Println(xi)

	y := []int {56, 57, 58}

	xi = append(xi, y...)

	// print
	fmt.Println(xi)
}
