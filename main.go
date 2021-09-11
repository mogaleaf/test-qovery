package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main(){
	http.Handle("/foo", &myHandler{})

	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

type myHandler struct{

}

func(*myHandler) ServeHTTP(res http.ResponseWriter,req *http.Request){
	log.Println("received")
	res.Write([]byte("received"))
}
