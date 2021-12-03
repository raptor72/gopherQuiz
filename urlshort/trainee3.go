package main

import (
//        "io"
        "log"
        "fmt"
        "net/http"
)

func main() {
        mux := defaultMux()
        fmt.Printf("%T\n", mux) // http.ServeMux -> http.HandlerFunc
        pathsToUrls := map[string]string{
                "/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
                "/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
        }
        mapHandler := MapHandler(pathsToUrls, mux)
        fmt.Println("start serve 8080")
        log.Fatal(http.ListenAndServe(":8080", mapHandler))
}

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
    redirect := func (w http.ResponseWriter, r *http.Request) {
        if val, ok := pathsToUrls[r.URL.Path]; ok {
            http.Redirect(w, r, val, 301)
        }
    }
    return redirect
}


func defaultMux() *http.ServeMux {
        mux := http.NewServeMux()
        mux.HandleFunc("/", hello)
        return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Hello, world!")
}
