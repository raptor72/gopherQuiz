package main

import (
        "io"
        "fmt"
        "net/http"
//        "github.com/gophercises/urlshort"
)

func main() {
        mux := defaultMux()
        pathsToUrls := map[string]string{
                "/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
                "/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
        }
        mapHandler := MapHandler(pathsToUrls, mux)


        fmt.Println("start serve 8080")
        http.ListenAndServe(":8080", mapHandler)
}


func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
    fmt.Println(pathsToUrls)
    fmt.Printf("%T\n", fallback)
    isHandler := func(w http.ResponseWriter, r *http.Request) {
        fmt.Println(r.URL.Path)
        io.WriteString(w, "Hello from a HandleFunc #1!\n")
        if val, ok := pathsToUrls[r.URL.Path]; ok {
            fmt.Println("is ok")
//            http.RedirectHandler(val, 307)
//            http.Handle(r.URL.Path, http.RedirectHandler(val, 307))
//            w.WriteHeader(http.StatusOK)
            http.Redirect(w, r, val, 307)
//            return
        } else {
            fmt.Println("is not ok")
            http.NotFound(w, r)
        }
    }
    fmt.Printf("%T\n", isHandler)
    return isHandler
}



func defaultMux() *http.ServeMux {
        mux := http.NewServeMux()
        mux.HandleFunc("/", hello)
        return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Hello, world!")
}
