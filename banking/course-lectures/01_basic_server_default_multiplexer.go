package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// define routes
	http.HandleFunc("/greet", greet)

	//starting server
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!!")
}
