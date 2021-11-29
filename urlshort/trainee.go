package main

import (
    "io"
    "net/http"
    "fmt"
)


func hello2(w http.ResponseWriter, req *http.Request) {
    w.Header().Set("Trailer", "AtEnd1, AtEnd2")
    w.Header().Add("Trailer", "AtEnd3")

    w.Header().Set("Content-Type", "text/plain; charset=utf-8") // normal header
    w.WriteHeader(http.StatusOK)

    w.Header().Set("AtEnd1", "value 1")
    io.WriteString(w, "This HTTP response has both headers before this text and trailers at the end.\n")
    w.Header().Set("AtEnd2", "value 2")
    w.Header().Set("AtEnd3", "value 3")
    fmt.Println("################################")
    fmt.Println(w)
    fmt.Println("################################")
//    fmt.Printf("%T", w) // *http.response
}


//func Redirect(w ResponseWriter, r *Request, url string, code int)
//func RedirectHandler(url string, code int) Handler
//func hello3()
//https://pkg.go.dev/net/http#RedirectHandler

func main() {
pathsToUrls := map[string]string{
    "/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
    "/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
}
fmt.Println(pathsToUrls)
fmt.Printf("%T\n", pathsToUrls)

h1 := func(w http.ResponseWriter, _ *http.Request) {
    io.WriteString(w, "Hello from a HandleFunc #1!\n")
}
h2 := func(w http.ResponseWriter, _ *http.Request) {
    io.WriteString(w, "Hello from a HandleFunc #2!\n")
}
mux := http.NewServeMux()
mux.HandleFunc("/send", hello2)
mux.HandleFunc("/", h1)
mux.HandleFunc("/endpoint", h2)

fmt.Println("Starting the trainee server on :8081")
http.ListenAndServe(":8081", mux)
}




