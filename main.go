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
	"time"
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
	timePtr := flag.Int("time-limit", 30, "the time limit (in seconds) to finish the whole quiz")
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
	var score int
	timer := time.NewTimer(time.Duration(*timePtr) * time.Second)

questionsLoop:
	for index, question := range questionList {
		fmt.Printf("Problem #%d: %s\n", index+1, question.question)

		firstChan := make(chan string)

		go func() {
			input, _ := reader.ReadString('\n')
			firstChan <- strings.TrimRight(input, "\n")
		}()

		select {
		case <-timer.C:
			break questionsLoop
		case in := <-firstChan:
			if in == question.answer {
				score++
			}
		}

	}
	fmt.Println("Your score is:", score, "/", len(questionList))
}
