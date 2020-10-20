package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, Greeting("Code.education Rocks! "))
	})
	http.ListenAndServe("0.0.0.0:80", nil)
}

func Greeting(text string) string {
	return "<b>" + text + "</b>"
}
