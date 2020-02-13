package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world"))
	})

	log.Println("Server listen on port 8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
