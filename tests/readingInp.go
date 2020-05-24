/*
 * Created on Sat May 16 2020
 *
 * Primary Author - Achin Gupta
 */

package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	var f *os.File
	f = os.Stdin
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(">", scanner.Text())
	}
}
