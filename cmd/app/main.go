package main

import (
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.RemoteAddr)
	log.Println("hello request2", r.RemoteAddr)
}

func main() {
	http.HandleFunc("/hello", hello)
	log.Println("server started")
	http.ListenAndServe(":4500", nil)
}
