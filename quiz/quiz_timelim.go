package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

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

func timeTimer(qzs []Quiz, timeout int) {
	timer := time.After(time.Second * time.Duration(timeout))
	<-timer
	printResult(qzs)
	os.Exit(0)
}

func run(qzs []Quiz, timeout int) {
	go timeTimer(qzs, timeout)
	var userAnswer int
	for idx, quiz := range qzs {
		fmt.Println(quiz.question)
		fmt.Scan(&userAnswer)
		quiz.user_answer = userAnswer
		qzs[idx] = quiz
	}
	printResult(qzs)
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
	filename := flag.String("filename", "problems.csv", "a string filename to load quiz")
	timeout := flag.Int("timeout", 4, "an int default timeout")
	flag.Parse()
	quizs := loadQuiz(*filename)
	run(quizs, *timeout)
}
