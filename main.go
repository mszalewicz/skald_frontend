package main

import (
	"fmt"
	"log"
	"net/http"

	"skald_frontend/templates"

	"github.com/a-h/templ"
)

func main() {
	// specifying 127.0.0.1 directly, overcomes problem with constant firewall warning when using Air hot reload
	const addr = "127.0.0.1:8080"

	fs := http.FileServer(http.Dir("static"))

	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.Handle("/", templ.Handler(templates.MainPage()))

	fmt.Println("Starting server on", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err.Error())
	}
}
