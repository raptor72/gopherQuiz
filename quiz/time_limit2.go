package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
        "time"
)

const readTimeOut = 4

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

/*
func timeTimer() {
    timer := time.After(time.Second * 4)
    <-timer
    fmt.Println("Time limit exceed")
    os.Exit(1)
}
*/

/*
func timeTimer() {
    timer := time.After(time.Second * 4)
    <-timer
    fmt.Println("Time limit exceed")
    exitChannel := make(chan int)
//    exitChannel <-1
    return exitChannel
}
*/


func timeTimer(r, out chan int) {
    for {
        timeout := time.After(time.Second * 3)
        select {
        case <-r:
            continue
        case <- timeout:
            data := <- r
            out <- data
            return
        }
    }
}



/*
func calculator(c, out chan int, done <-chan time.Time) chan int {
//    out := make(chan int)
    var sum int
    go func() {
//	defer close(out)
	for {
	    select {
	    case data := <-c:
		sum += data
	    case <-done:
		out <- sum
                return
            }
	}
    }()
    return out
}
*/


func main() {
        c := make(chan int, 12)
//        re := make(chan int)
//        quit := make(chan int)
        out := make(chan int)
//        timer := time.After(time.Second * 2)
//        fmt.Printf("%T", timer)
        go timeTimer(c, out)
//        res := calculator(c, timer)
//        fmt.Printf("%T, %T", start, readTimeOut)
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
                c <- count
        }
	fmt.Printf("Total questions is: %v, right questions is: %v\n", count, rights)
}
