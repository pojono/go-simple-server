package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	server := http.FileServer(http.Dir("./static"))
	http.Handle("/", server)
	http.HandleFunc("/form", formHandler)
	fmt.Printf("Server starting!")
	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatalf("Server error: %s\n", err)
	}
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Parse error: %v", err)
		return
	}
	name := r.FormValue("name")
	fmt.Fprintf(w, "Your name is %v", name)
}
