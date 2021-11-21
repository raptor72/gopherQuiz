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
func timeTimer(exitChannel chan int) {
    timer := time.After(time.Second * 4)
    <-timer
    fmt.Println("Time limit exceed")
    exitChannel <- 1
}
*/

func timeTimer(c, r, quit chan int) {
    defer close(c)
    defer close(r)
    for {
        select {
        case <-c:
//            count := <-c
            continue
//            fmt.Println(count)
        case <-r:
//            rights := <-r
            continue
        case <-time.After(time.Second * 3):
            count, rights := <-c, <-r
            fmt.Println(count, rights)
            fmt.Println("Time limit exceed")
//            os.Exit(1)
//            return
            quit <- 1
        }
    }
}


func main() {
        c := make(chan int)
        re := make(chan int)
        quit := make(chan int)
//        timer := time.After(time.Second * 2)
//        fmt.Printf("%T", timer)
        go timeTimer(c, re, quit)
//        go timeTimer(c)
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
                c <- count
                re <- rights
                select {
                case <- quit:
                    fmt.Printf("Total questions is: %v, right questions is: %v\n", count, rights)
                    os.Exit(0)
                }

        }

/*
        go func() {
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
                re <- rights
                select {
                case <- quit:
                    fmt.Printf("Total questions is: %v, right questions is: %v\n", count, rights)
                }
            }
        }()
*/
	fmt.Printf("Total questions is: %v, right questions is: %v\n", count, rights)
}
