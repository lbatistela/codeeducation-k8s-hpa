package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", SquareRootServer)
	http.ListenAndServe(":8000", nil)
}

func SquareRootServer(w http.ResponseWriter, r *http.Request) {
	sqrtNTimes(0.0001, 100000000)
	fmt.Fprintf(w, "Code.education Rocks!")
}