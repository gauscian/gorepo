/*
Hands-on exercise #3
	Using goroutines, create an incrementer program
		have a variable to hold the incrementer value
		launch a bunch of goroutines
			each goroutine should
				read the incrementer value
					store it in a new variable
				yield the processor with runtime.Gosched()
				increment the new variable
				write the value in the new variable back to the incrementer variable
	use waitgroups to wait for all of your goroutines to finish
	the above will create a race condition.
	Prove that it is a race condition by using the -race flag
	if you need help, here is a hint: https://play.golang.org/p/FYGoflKQej
*/

package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	counter := 0
	counter2 := 0
	fmt.Println("Number of cpus = ", runtime.NumCPU())

	// go func run counter
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			v := counter
			v1 := counter2
			v++
			v1++
			counter = v
			counter2 = v1
			wg.Done()
		}()
		fmt.Println("Still adding # routines = ", runtime.NumGoroutine())
	}
	fmt.Println("num of go routines = ", runtime.NumGoroutine())
	fmt.Println("counter = ", counter)
	fmt.Println("counter2 = ", counter2)
	wg.Wait()
}
