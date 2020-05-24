/*
Use Mutex to remove race
verify with race flag.
*/

package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup
	var counter int64
	fmt.Println("Number of cpus = ", runtime.NumCPU())
	wg.Add(100)

	// go func run counter
	for i := 0; i < 100; i++ {
		go func(i int) {
			fmt.Println("Current func ", i)
			atomic.AddInt64(&counter, 1)
			wg.Done()
		}(i)
	}
	// placement of wait group matters alot.
	// Trying to read before waiting will be problemmatic because you will
	// be reading a race condition variable.
	wg.Wait()
	fmt.Println("num of go routines after wg dones = ", runtime.NumGoroutine())
	fmt.Println("counter = ", counter)
}
