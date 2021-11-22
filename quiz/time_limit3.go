package main

import (
    "time"
    "os"
    "encoding/csv"
    "log"
    "fmt"
    "strconv"
)

type Quiz struct {
    question string
    right_answer int
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


func timeTimer(qzs []Quiz) {
    timer := time.After(time.Second * 4)
    <-timer
    printResult(qzs)
    os.Exit(0)
}


func run(qzs []Quiz) []Quiz {
    go timeTimer(qzs)
    var userAnswer int
    for idx, quiz := range qzs {
        fmt.Println(quiz.question)
        fmt.Scan(&userAnswer)
        quiz.user_answer = userAnswer
        qzs[idx] = quiz
    }
    return qzs
}


func printResult(qzs []Quiz) {
    total_count := len(qzs)
    var right_answers int
    for _, quiz := range qzs {
        if quiz.right_answer == quiz.user_answer {
            right_answers++
        }
    }
    fmt.Printf("Total questions: %d right answers: %d\n", total_count, right_answers)
}


func main() {
    quizs := load_quiz("problems.csv")
    run(quizs)
}

