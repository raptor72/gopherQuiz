package urlshort

import (
//        "bytes"
//        "io"
        "fmt"
//        "os"
        "testing"
//        "io/ioutil"
)


func TestBytesToString(t *testing.T) {
        bts := []byte("dqwdqwd")
        expected_string := "dqwdqwd"
        st := bytesToString(bts)
        if st != expected_string {
            t.Fatalf("Want %v\n, but got %v", expected_string, st)
        } else {
            fmt.Printf("Want `%v`, and got `%v`. Ok. ", expected_string, st)
        }
}


