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


func bytesToString(data []byte) string {
        return string(data[:])
}


func parseYaml(yml []byte) []yamlStruct {
        var yamlArray []yamlStruct
        var y yamlStruct
        stringYaml := bytesToString(yml)
        s := strings.Split(stringYaml, "- ")
        for _, yamlBlock := range s {
                yamlBlock = strings.ReplaceAll(yamlBlock, "  ", "")
                if len(yamlBlock) == 0 {
                    continue
                }
                err := yaml.Unmarshal([]byte(yamlBlock), &y)
                if err != nil {
                        fmt.Println(err)
                }
                if y.URL != "" {
                        yamlArray = append(yamlArray, y)
                }
        }
        return yamlArray
}


func buildMap(yamlArray []yamlStruct) map[string]string {
        result := make(map[string]string)
        for _, yamlStruct := range yamlArray {
                result[yamlStruct.PATH] = yamlStruct.URL
        }
        return result
}


func main() {
        yml := `
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`
        byaml := []byte(yml)
//        fmt.Println(string(byaml))
//        fmt.Println(bytesToString(byaml))
        parsedYaml := parseYaml(byaml)
        mapArray := buildMap(parsedYaml)
        fmt.Println(mapArray)
}
