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



//func (quiz *Quiz) answered(ans int) {
//    quiz.user_answer = ans
//}


func run(qzs []Quiz) []Quiz {
// [{5+5 10 0} {1+1 2 0} {8+3 11 0} {1+2 3 0} {8+6 14 0} {3+1 4 0} {1+4 5 0} {5+1 6 0} {2+3 5 0} {3+3 6 0} {2+4 6 0} {5+2 7 0}]
    count := len(qzs)
    var userAnswer int
    fmt.Printf("Total count of questions is: %d\n", count)
    for idx, quiz := range qzs {
        fmt.Println(quiz)
        fmt.Scan(&userAnswer)
        quiz.user_answer = userAnswer
        qzs[idx] = quiz
    }
    return qzs
}

func main() {
    quizs := load_quiz("problems.csv")
    fmt.Println(quizs) // [{5+5 10 0} {1+1 2 0} ...]
    fmt.Println(run(quizs))
}

