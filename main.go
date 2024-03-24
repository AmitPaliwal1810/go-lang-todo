package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "mojo-jojo")
	})

	err := http.ListenAndServe(":5000", nil)

	if err != nil {
		fmt.Println("Error starting server", err)
	}
}
