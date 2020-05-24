/*
Use Atomic to remove race conditions
verify with race flag.
*/

package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var rwmu sync.Mutex
	counter := 0
	fmt.Println("Number of cpus = ", runtime.NumCPU())
	wg.Add(100)

	// go func run counter
	for i := 0; i < 100; i++ {
		go func() {
			rwmu.Lock()
			v := counter
			v++
			counter = v
			rwmu.Unlock()
			wg.Done()
		}()
	}
	// placement of wait group matters alot.
	// Trying to read before waiting will be problemmatic because you will
	// be reading a race condition variable.
	wg.Wait()
	fmt.Println("num of go routines = ", runtime.NumGoroutine())
	fmt.Println("counter = ", counter)
}
