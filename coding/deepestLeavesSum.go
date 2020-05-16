package main

import "fmt"


func main() {
	sl := []int{}
	sl = append(sl, 3, 3, 4, 5, 5, 6, 7)
	ele, sl := sl[0], sl[1:] 
	fmt.Println(ele, "-->" ,sl)
	ele, sl = sl[0], sl[1:] 
	fmt.Println(ele, "-->" ,sl)
	// for len(sl) > 0 {
	// 	fmt.Println(len(sl))
	// 	ele := 0
	// 	ele, sl = sl[0], sl[1:] 
	// 	fmt.Println(ele, sl[0])
	// }
}