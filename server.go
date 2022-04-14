package main

import (
  "net/http"
  "fmt"
  "os"
)

func main() {
	http.HandleFunc("/", Hello)
	http.ListenAndServe(":8080", nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {
  name := os.Getenv("NAME")
  age := os.Getenv("AGE")

  fmt.Fprintf(w, "Hello %s! You are %s years old.", name, age)
	w.Write([]byte("<h1>Hello Full Cycle!</h1>"))
}
