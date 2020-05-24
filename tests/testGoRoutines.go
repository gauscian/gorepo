/*
* Created on Sun May 17 2020
Testing out go routines
* Primary Author - Achin Gupta
* Copyright (c) 2020 Netskope
*/

package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("os = ", runtime.GOOS)
	fmt.Println("number of cpus = ", runtime.NumCPU())
	fmt.Println("number of go routines = ", runtime.NumGoroutine())
	go foo()
	time.Sleep(1 * time.Second)
	fmt.Println("number of cpus = ", runtime.NumCPU())
	fmt.Println("number of go routines = ", runtime.NumGoroutine())
	bar()
}

func foo() {
	for i := 0; i <= 10; i++ {
		fmt.Println("foo = ", i)
	}
}

func bar() {
	for i := 0; i <= 10; i++ {
		fmt.Println("bar = ", i)
	}
}
