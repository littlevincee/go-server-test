package main

import (
	"fmt"
	"log"
	"net/http"
)

func seatingPlanHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/seating-plan" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	fmt.Fprintf(w, "Hi")
}

func main() {
	htmlFileDir := "./static"

	fileServer := http.FileServer(http.Dir(htmlFileDir))

	http.Handle("/", fileServer)

	http.HandleFunc("/seating-plan", seatingPlanHandler)

	fmt.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
