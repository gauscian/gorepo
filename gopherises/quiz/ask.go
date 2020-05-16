package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"gophercises/quiz/code"
	"bufio"
)

// QInstance is one question instance
type QInstance code.Question


func main() {
	data := readCsv("problems.csv")
	askQuestions(data)
}

func readCsv(filePath string) ([]QInstance) {
	// open the file for reading
	f, err := os.Open(filePath)
	if err != nil {
		return []QInstance{}
	}

	// close the file after program exits
	defer f.Close()

	// read csv file and return data as a slice
	lines, _ := csv.NewReader(f).ReadAll()
	data := []QInstance{}
	for _, line := range lines {
		ques := QInstance{}
		ques.QuestionText = line[0]
		ques.AnswerText = line[1]
		data = append(data, ques)
	}
	return data
}


func askQuestions(data []QInstance) {
	scanner := bufio.NewScanner(os.Stdin)
	correct := 0
	incorrect := 0
	total := 0
	for _, d := range data {
		fmt.Print("Question: ", d.QuestionText, "\n", "Answer: ")
		// Scans a line from Stdin(Console)
		scanner.Scan()
		// Holds the string that scanned
		text := scanner.Text()
		if text == d.AnswerText {
			fmt.Println("You are right")
			correct++
		} else {
			fmt.Println("Incorrect Try again")
			incorrect++
		}
		total++
	}
	fmt.Println("Questions asked : ",total)
	fmt.Println("Questions correct : ",correct)
	fmt.Println("Questions incorrect : ",incorrect)
}
