package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

const defaultProblemFilename = "problems.csv"

var (
	correctAnswers int
	totalQuestions int
)

func startQuiz() chan bool {
	done := make(chan bool)
	go func() {
		f, err := os.Open(defaultProblemFilename)
		if err != nil {
			fmt.Printf("Error while opening the file %v\n", err)
			done <- false
		}
		defer f.Close()
		// fmt.Println("Hello from egypt")
		// read the file using csv package
		r := csv.NewReader(f)
		records, err := r.ReadAll()
		if err != nil {
			fmt.Printf("Error while reading the file %v\n", err)
			done <- false
		}
		totalQuestions = len(records)
		for index, record := range records {
			question, correctAnswer := record[0], record[1]
			var answer string
			fmt.Printf("%d. %s\n", index+1, question)
			if _, err := fmt.Scan(&answer); err != nil {
				fmt.Printf("failed to scan %v\n", err)
			}
			if answer == correctAnswer {
				correctAnswers++
			}
		}
		done <- true
	}()

	return done
}
func main() {
	quizDone := startQuiz()
	quizTimeDone := time.NewTimer(3 * time.Second).C

	select {
	case <-quizDone:
	case <-quizTimeDone:
	}
	fmt.Printf("Result %d/%d\n", correctAnswers, totalQuestions)
}
