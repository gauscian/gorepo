package main

import "fmt"

func main() {
	fmt.Println(numberOfSteps(6))
}
func numberOfSteps(num int) int {
	return getSteps(num)
}

func getSteps(num int) int {
	if num == 0 {
		return 0
	} else if num%2 == 0 {
		num /= 2
	} else {
		num -= 1
	}

	return 1 + getSteps(num)
}
