package main

import (
    "fmt"
    "log"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello world %s! (%s)", r.URL.Path[1:],"v1.0")
}

func main() {
    http.HandleFunc("/", handler)
    log.Print("Starting on 8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
