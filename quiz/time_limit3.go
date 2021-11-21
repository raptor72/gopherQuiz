package main

import (
    "os"
    "encoding/csv"
    "log"
    "fmt"
    "strconv"
)

type Quiz struct {
    question string
    answer int
    user_answer int
}

func load_quiz(csv_file string) []Quiz {
        var quizs []Quiz
        file, err := os.Open(csv_file)
        if err != nil {
                log.Fatal(err)
        }
        defer file.Close()
        r := csv.NewReader(file)
        rows, err := r.ReadAll()
        if err != nil {
            log.Fatal(err)
        }
        for _, row := range rows {
            if len(row) != 2 {
                    log.Fatal("Wrong row length")
            }
            panswer, err := strconv.Atoi(row[1])
            if err != nil {
                    log.Fatal("error of type result")
            }

            quiz := Quiz{row[0], panswer, 0}
            quizs = append(quizs, quiz)
        }
        return quizs
}

func main() {
    quizs := load_quiz("problems.csv")
    fmt.Println(quizs) // [{5+5 10 0} {1+1 2 0} ...]
}

