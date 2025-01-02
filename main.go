package main

import (
	"fmt"
	"net/http"
)

func main() {
	const addr = "127.0.0.1:8080"

	http.HandleFunc("/", mainPage)

	http.ListenAndServe(addr, nil)
}

func mainPage(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "main page, yes or not or yes")
}