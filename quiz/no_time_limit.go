package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func readQuestions(row []string) string {
	if len(row) != 2 {
		log.Fatal("Wrong row length")
	}
	return row[0]
}

func parseAnswer(row []string) int {
	if len(row) != 2 {
		log.Fatal("Wrong row length")
	}
	result, err := strconv.Atoi(row[1])
	if err != nil {
		log.Fatal("error of type result")
	}
	return result
}

func main() {
	var count, rights, userAnswer int
	file, err := os.Open("problems.csv")
	//    file, err := os.Open("test.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	r := csv.NewReader(file)

	for {
		row, err := r.Read()
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
		if err == io.EOF {
		//       fmt.Println("read all file")
			break
		}
		if len(row) != 2 {
			log.Fatalf("Uncoreect row length. Should be 2 but found: %v", len(row))
		}
		fmt.Println(readQuestions(row))
		fmt.Scan(&userAnswer)
		if parseAnswer(row) == userAnswer {
			rights++
		}
		count++
	}
	fmt.Printf("Total questions is: %v, right questions is: %v\n", count, rights)
}
