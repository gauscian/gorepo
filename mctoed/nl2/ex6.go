package main

import "fmt"

func main() {
	const (
		a = 2017 + iota
		b = iota
		c = iota
		d = iota
	)

	fmt.Printf("%d, %d, %d, %d", a, b, c, d)
}