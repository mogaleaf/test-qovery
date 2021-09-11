package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/service1", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "you talk to service 1")
	})

	http.HandleFunc("/service2", func(w http.ResponseWriter, r *http.Request) {
		resp, err := http.Get("http://localhost:8082/service2")
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		fmt.Fprintf(w, string(body))
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}
