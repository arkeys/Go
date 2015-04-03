package main

import (
	"io"
	"net/http"
)

func welc(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Welcome to go!")
}

func main() {
	http.HandleFunc("/", welc)
	http.ListenAndServe(":80", nil)
}
