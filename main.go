package main

import (
	"RyanFin/GoQuizApp/pkg/constants"
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	File string `yaml:"file"`
}

func main() {
	// load config file from scratch
	filename, err := filepath.Abs(constants.FileName)
	if err != nil {
		panic(err)
	}

	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	// remember to close the file at the end of the program
	defer f.Close()

	// load quiz data from csv file
	quizData, err := loadCSVToSlice(f)
	if err != nil {
		panic(err)
	}

	fmt.Println(quizData)

	// start quiz
	startQuiz(quizData)
}

func loadCSVToSlice(f *os.File) ([][]string, error) {
	r := csv.NewReader(f)

	problems, err := r.ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	return problems, nil
}

func startQuiz(quizData [][]string) {
	score := 0
	maxScore := len(quizData)

	// start quiz
	for i, e := range quizData {
		question := e[0]
		answer := e[1]

		fmt.Printf("\n[Question %v] What is %s:\n", i, question)
		fmt.Print("Enter your answer: ")

		var input string
		fmt.Scanln(&input)

		if input != answer {
			fmt.Println("Incorrect.")
			continue
		}

		if input == answer {
			fmt.Println("Correct :)")
			score += 1
		}
	}

	fmt.Printf("\nFinal Score: %v/%v\n", score, maxScore)
}
