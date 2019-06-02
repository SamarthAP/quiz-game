package main

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
)

type Question struct {
	question string
	answer   string
}

func main() {
	csvFile, err := os.Open("quiz.csv")
	if err != nil {
		log.Fatal(err)
	}
	reader := csv.NewReader(bufio.NewReader(csvFile))
	questionList := make([]Question, 0)
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		questionList = append(questionList, Question{
			question: line[0],
			answer:   line[1],
		})
	}
}
