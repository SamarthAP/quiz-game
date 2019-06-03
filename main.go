package main

import (
	"bufio"
	"encoding/csv"
	"flag"
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
	filePtr := flag.String(
		"csv",
		"quiz.csv",
		"a csv file with the format question,answer (default: quiz.csv)")
	flag.Parse()

	csvFile, err := os.Open(*filePtr)
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
