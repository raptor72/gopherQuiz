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

func parseYaml(stringYaml string) []yamlStruct {
        var yamlArray []yamlStruct
        var y yamlStruct
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
        parsedYaml := parseYaml(yml)
        mapArray := buildMap(parsedYaml)
        fmt.Println(mapArray)
}
