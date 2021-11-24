package main

import (
	"bytes"
	"io"
	"os"
	"testing"
	"io/ioutil"
)

func TestLoadQuis(t *testing.T) {
	file := "problems.csv"
	expected_qzs := []Quiz{Quiz{"5+5", 10, 0}, Quiz{"1+1", 2, 0}, Quiz{"8+3", 11, 0}, Quiz{"1+2", 3, 0}, Quiz{"8+6", 14, 0}, Quiz{"3+1", 4, 0}, Quiz{"1+4", 5, 0}, Quiz{"5+1", 6, 0}, Quiz{"2+3", 5, 0}, Quiz{"3+3", 6, 0}, Quiz{"2+4", 6, 0}, Quiz{"5+2", 7, 0}}
	qzs := loadQuiz(file)
	for idx, quiz := range qzs {
		if quiz != expected_qzs[idx] {
			t.Fatalf("Want %v\n, but got %v", expected_qzs, qzs)
		}
	}
}

func TestPrintResult(t *testing.T) {
	quiz := []Quiz{Quiz{"5+5", 10, 10}, Quiz{"1+1", 2, 2}, Quiz{"8+3", 11, 0}}
	var expected_string string = "Total questions: 3 right answers: 2\n"
	old := os.Stdout // keep backup of the real stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	printResult(quiz)

	outC := make(chan string)
	// copy the output in a separate goroutine so printing can't block indefinitely
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()
	// back to normal state
	w.Close()
	os.Stdout = old // restoring the real stdout
	out := <-outC
	// reading our temp stdout
	//    fmt.Println(out)
	if out != expected_string {
		t.Fatalf("Want %v\n, but got %v", expected_string, out)
	}
}

func TestPrintResult2(t *testing.T) {
	quiz := []Quiz{Quiz{"5+5", 10, 10}, Quiz{"1+1", 2, 2}, Quiz{"8+3", 11, 0}}
	var expected_string string = "Total questions: 3 right answers: 2\n"
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	printResult(quiz)
	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout
	if string(out) != expected_string {
		t.Fatalf("Want %v\n, but got %v", expected_string, string(out))
	}
}
