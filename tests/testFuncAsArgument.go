/*
 * Created on Sat May 16 2020
 *
 * Primary Author - Achin Gupta
 */

package main

import (
	"fmt"
)

func main() {
	allInts := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	allSum := sum(allInts...)
	fmt.Println("All ints sum ", allSum)

	// evenSum var and evenSum has the same name, go
	// was able to resolve what is what.
	evenSum := evenSum(sum, allInts...)
	fmt.Println("even Sum", evenSum)
}

// here is a func of format (...int) int
func sum(ii ...int) int {
	total := 0
	for _, v := range ii {
		total += v
	}
	return total
}

// separates out the even number and sums then using the operator method.
func evenSum(sum func(ii ...int) int, allInts ...int) int {
	ei := []int{}
	for _, v := range allInts {
		if v%2 == 0 {
			ei = append(ei, v)
		}
	}
	return sum(ei...)
}
