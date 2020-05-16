package main

import "fmt"

func main() {
	i := 1991
	for {
		fmt.Printf("I was alieve in %d\n", i)
		if i > 2020 {
			break
		}
		i++
	}
}