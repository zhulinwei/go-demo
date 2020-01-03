package main

import "fmt"
import "net/http"

func hello (w http.ResponseWriter, req *http.Request) {
  fmt.Fprintf(w, "hello\n")
}

func headers (w http.ResponseWriter, req *http.Request) {
  for name, headers := range req.Header {
    for _, header := range headers {
      fmt.Fprintf(w, "%v: %v\n", name, header)
    }
  }
}

func main () {
  http.HandleFunc("/hello", hello)
  http.HandleFunc("/headers", headers)

  http.ListenAndServe(":3000", nil)
}
