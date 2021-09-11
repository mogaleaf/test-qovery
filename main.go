package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/service1", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "you talk to service 1")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
