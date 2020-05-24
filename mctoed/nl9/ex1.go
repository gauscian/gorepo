/*
in addition to the main goroutine, launch two additional goroutines
	each additional goroutine should print something out
use waitgroups to make sure each goroutine finishes before your program exists
*/
// author Achin Gupta

package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	fmt.Println("Number of CPUs", runtime.NumCPU())

	// How many wait groups do you want to add
	wg.Add(2)
	// go routine 2
	go func() {
		fmt.Println("go routine 2")
		wg.Done()
	}()

	// go routine 1
	go func() {
		fmt.Println("go routine 1")
		wg.Done()
	}()

	fmt.Println("main routine")
	// waits until the wait group counter is 0
	wg.Wait()
}
