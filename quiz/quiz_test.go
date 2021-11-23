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

func loadQuiz(csv_file string) []Quiz {
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
*/

/*
func TestprintResult(t *testing.T) {
    a := Quiz{"1+1",2, 2}
    b := Quiz{"1+2",3, 3}
    c := Quiz{"3+8",11, 11}
    qzs := []Quiz{a,b,c}
    testingResult := printResult(qzs)
    rigthResult := "Total questions: 3 right answers: 3\n"
    if testingResult != rigthResult {
        t.Fatalf("Want %v, but got %v", rigthResult, testingResult)
    }
}
*/

func TestLoadQuis(t *testing.T) {
    file := "problems.csv"
    expected_qzs := []Quiz{Quiz{"5+5", 10, 0}, Quiz{"1+1", 2, 0}, Quiz{"8+3", 11, 0}, Quiz{"1+2", 3, 0}, Quiz{"8+6", 14, 0}, Quiz{"3+1", 4, 0}, Quiz{"1+4", 5, 0}, Quiz{"5+1", 6, 0}, Quiz{"2+3", 5, 0}, Quiz{"3+3", 6, 0}, Quiz{"2+4", 6, 0}, Quiz{"5+2", 7, 0}}
    qzs := loadQuiz(file)
    for idx, quiz := range(qzs) {
        if quiz != expected_qzs[idx] {
            t.Fatalf("Want %v\n, but got %v", expected_qzs, qzs)
        }
    }
}




