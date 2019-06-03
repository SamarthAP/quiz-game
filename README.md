# quiz-game

A Go program that will read in a quiz from a CSV file and ask questions that the user is expected to try and answer. After all the questions have been asked, the program displays the number of correct answers.

## Show flags
```
$ go build .
$ ./quiz-game -h
Usage of ./quiz-game:
  -csv string
        a csv file with the format question,answer (default "quiz.csv")
```

## Run
```
go build .
./quiz-game
```