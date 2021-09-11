package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.Handle("/foo", &myHandler{})

	http.HandleFunc("/service1", func(w http.ResponseWriter, r *http.Request) {
		resp, err := http.Get("http://localhost:8081/service1")
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		fmt.Fprintf(w, string(body))
	})
	http.HandleFunc("/service2", func(w http.ResponseWriter, r *http.Request) {
		resp, err := http.Get("http://localhost:8081/service2")
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		fmt.Fprintf(w, string(body))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

type myHandler struct {
}

func (*myHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	log.Println("received")
	res.Write([]byte("received"))
}
