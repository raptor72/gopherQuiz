package main

import (
    "fmt"
    "strings"
    "gopkg.in/yaml.v2"
)

func main() {
        pathsToUrls := map[string]string{
                "/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
                "/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
        }

        fmt.Sprintln(pathsToUrls)

        yml := `
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`

yml2 := `
path: /urlshort
url: https://github.com/gophercises/urlshort
`

type yamlStruct struct {
    PATH string `yaml:"path"`
    URL string `yaml:"url"`
}

var y yamlStruct
var y2 yamlStruct
fmt.Sprintln(y)
fmt.Sprintln(y2)


err := yaml.Unmarshal([]byte(yml2), &y2)
if err != nil {
    fmt.Println(err)
}
fmt.Println(y2)


s := strings.Split(yml, "- ")

for _, val := range s {

//    val = strings.TrimSpace(val)
    val = strings.ReplaceAll(val, "  ", "")
//    text = strings.ToLower(text)
//    val = strings.Trim(val, "\n")


    fmt.Println(val)

//    fmt.Println(len(val))

    if len(val) == 0 {
        continue
    }


    err := yaml.Unmarshal([]byte(val), &y)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(y)
    fmt.Printf("%T", y)

}

fmt.Printf("%T", s)


/*
err := yaml.Unmarshal([]byte(yml), &y)
if err != nil {
    log.Fatalf("cannot unmarshal data: %v", err)
}
fmt.Println(y)
fmt.Println(Y)

type T struct {
    F string `yaml:"a,omitempty"`
    B int
}
var t T
yaml.Unmarshal([]byte("a: efefe\nb: 2"), &t)

fmt.Println(t.F)
fmt.Print(t)
*/
//        yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), mapHandler)


}