package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/service2", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "you talk to service 2")
	})

	log.Fatal(http.ListenAndServe(":8082", nil))
}
