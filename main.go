package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
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
	csvReader := csv.NewReader(bufio.NewReader(csvFile))
	questionList := make([]Question, 0)
	for {
		line, error := csvReader.Read()
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
	reader := bufio.NewReader(os.Stdin)
	var score [2]int // correct, total
	for _, question := range questionList {
		fmt.Println(question.question)
		input, _ := reader.ReadString('\n')
		if strings.TrimRight(input, "\n") == question.answer {
			score[0] = score[0] + 1
		}
		score[1] = score[1] + 1
	}
	fmt.Println("Your score is:", score[0], "/", score[1])
}
