package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"strings"
)

type yamlStruct struct {
	PATH string `yaml:"path"`
	URL  string `yaml:"url"`
}


/*
//func buildMap(parsedYaml []yamlStruct) { //map[string]string {
func buildMap(parsedYaml []yamlStruct) {
//	var result map[string]string

	for val := range parsedYaml {
		fmt.Println(val)
	}
}
*/

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
	type yamlStruct struct {
		PATH string `yaml:"path"`
		URL  string `yaml:"url"`
	}
	var yamlArray []yamlStruct
	var y yamlStruct

	s := strings.Split(yml, "- ")

	for _, val := range s {
		//    val = strings.TrimSpace(val)
		val = strings.ReplaceAll(val, "  ", "")
		//    text = strings.ToLower(text)
		//    val = strings.Trim(val, "\n")
		//    fmt.Println(val)
		if len(val) == 0 {
			continue
		}
		err := yaml.Unmarshal([]byte(val), &y)
		if err != nil {
			fmt.Println(err)
		}
		if y.URL != "" {
			yamlArray = append(yamlArray, y)
		}
	}
	fmt.Println(yamlArray)
//        fmt.Printf("%T\n", yamlArray)
        result := make(map[string]string)
	for _, val := range yamlArray {
//		fmt.Println(val)
                result[val.PATH] = val.URL
	}
        fmt.Println(result)

//	buildMap(yamlArray)


}
