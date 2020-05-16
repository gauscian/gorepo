package main

import (
	"flag"
	"fmt"
)

func main() {
	s, i := readArguments()
	fmt.Println(s, i)
}

func readArguments() (string, int) {
	filename := flag.String("filename", "problem.csv", "CSV File that conatins quiz questions")
	timeLimit := flag.Int("limit", 30, "Time Limit for each question")
	flag.Parse()
	return *filename, *timeLimit
}