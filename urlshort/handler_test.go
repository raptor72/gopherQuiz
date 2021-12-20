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
            fmt.Println("TestBytesToString: ")
            fmt.Printf("Want string `%v`, and got `%v`. Ok. \n\n", expected_string, st)
        }
}

func TestParseYAMLsingleBlock(t *testing.T) {
        single_yaml := []byte("- path: /urlshort\n  url: https://github.com/gophercises/urlshort")
        expected_yaml_struct := []yamlStruct{yamlStruct{"/urlshort", "https://github.com/gophercises/urlshort"} }
        ya, err := parseYAML(single_yaml)
        if err != nil {
            t.Fatalf("Get a parsing yaml error: %v", err)
        }
        for idx, ys := range(ya) {
            if ys.PATH != expected_yaml_struct[idx].PATH {
                t.Fatalf("Wrong parsing shortpath Want %v, but got %v", expected_yaml_struct[idx].PATH, ys.PATH)
            }
            if ys.URL != expected_yaml_struct[idx].URL {
                t.Fatalf("Wrong parsing full url. Want %v, but got %v", expected_yaml_struct[idx].URL, ys.URL)
            }
        }
        fmt.Println("TestParseYAMLsingleBlock: ")
        fmt.Printf("Got correct path and url from yaml in []bytes.\n\n")
}


func TestParseYAMLmultipleBlock(t *testing.T) {
        multiple_yaml := []byte("- path: /short_path1\n  url: https://url1\n- path: /short_path2\n  url: https://url2")
        expected_yaml_struct := []yamlStruct{yamlStruct{"/short_path1", "https://url1"}, yamlStruct{"/short_path2", "https://url2"} }
        ya, err := parseYAML(multiple_yaml)
        if err != nil {
            t.Fatalf("Get a parsing yaml error: %v", err)
        }
        for idx, ys := range(ya) {
            if ys.PATH != expected_yaml_struct[idx].PATH {
                t.Fatalf("Wrong parsing shortpath Want %v, but got %v", expected_yaml_struct[idx].PATH, ys.PATH)
            }
            if ys.URL != expected_yaml_struct[idx].URL {
                t.Fatalf("Wrong parsing full url. Want %v, but got %v", expected_yaml_struct[idx].URL, ys.URL)
            }
        }
        fmt.Println("TestParseYAMLmultipleBlock: ")
        fmt.Printf("Got correct path and url from yaml in []bytes.\n\n")
}
