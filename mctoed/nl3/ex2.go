package main

import "fmt"

func main() {
	for i := 65; i <= 90; i++ {
		fmt.Printf("%d", i)
		for j := 0; j <= 2; j++ {
			fmt.Printf("\n\t%#U\n", i)
		}
	}
}
