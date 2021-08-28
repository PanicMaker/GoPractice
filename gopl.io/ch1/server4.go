package main

import (
	"log"
	"net/http"
)

func main() {
	handler := func(writer http.ResponseWriter, request *http.Request) {
		lissajous(writer)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("0.0.0.0:8000", nil))
}
