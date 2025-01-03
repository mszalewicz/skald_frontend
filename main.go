package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"skald_frontend/templates"

	"github.com/a-h/templ"
)

func main() {
	// specifying 127.0.0.1 directly, overcomes problem with constant firewall warning when using Air hot reload
	const addr = "127.0.0.1:8080"

	fs := http.FileServer(http.Dir("static"))

	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.Handle("/", templ.Handler(templates.MainPage()))
	http.HandleFunc("/events", sseHandler)

	fmt.Println("Starting server on", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err.Error())
	}
}

func sseHandler(w http.ResponseWriter, r *http.Request) {
	// Set the necessary headers for SSE
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	// Ensure the connection stays open
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported", http.StatusInternalServerError)
		return
	}

	// Send events to the client in a loop
	for i := 0; i < 10; i++ {
		fmt.Fprintf(w, "event: newmsg\n")
		fmt.Fprintf(w, "data: <h3>Message: %d</h3>\n\n", i)
		flusher.Flush()             // Ensure data is sent to the client immediately
		time.Sleep(1 * time.Second) // Simulate a delay between messages
	}

	// Close the connection after sending all events
	// fmt.Fprintf(w, "event: newmsg\n")
	// fmt.Fprintf(w, "data: Closing connection\n\n")
	// flusher.Flush()
}
