package main

import (
    "testing"
)


/*
type Quiz struct {
        question     string
        right_answer int
        user_answer  int
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
*/

func TestprintResult(t *testing.T) {
    a := Quiz{"1+1",2, 2}
    b := Quiz{"1+2",3, 3}
    c := Quiz{"3+8",11, 11}
    qzs := [3]Quiz{a,b,c}
    testingResult := printResult(qzs)
    rigthResult := "Total questions: 3 right answers: 3\n"
    if testingResult != rigthResult {
        t.Fatalf("Want %v, but got %v", rigthResult, testingResult)
    }
}
