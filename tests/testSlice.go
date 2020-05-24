package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4}
	sliceme(&slice)
}

func sliceme(slice *[]int) {
	var_slice := (*slice)
	fmt.Println((var_slice)[1])
	fmt.Println(&var_slice)
	fmt.Println(slice)
	// newslice := append([]int{}, slice[index:index+1]...)
	// newslice = append(newslice, slice[0:index]...)
	// newslice = append(newslice, slice[index+1:]...)
	// fmt.Println(newslice)

}
