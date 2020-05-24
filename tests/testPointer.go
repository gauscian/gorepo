/*
* Created on Sun May 17 2020
*
* Primary Author - Achin Gupta
* Copyright (c) 2020 Netskope
*/

package main

import (
	"fmt"
)

func main() {
	var x int = 9
	// * here only works here on *T i.e pointer to T type.
	// * used with T and f* used with a variable are different stories.
	fmt.Println("Before mutation value for x = ", *&x)
	mutateInt(*&x)
	fmt.Println("after mutation x = ", x)

	var xi []int = []int{1,2}
	fmt.Println("Before ", xi)
	mutateSlice(&xi)
	fmt.Println(xi)

}

func mutateSlice(sl *[]int) {
	*sl = append(*sl, 3)
}

func mutateInt(i int) {
	i = -9
}
