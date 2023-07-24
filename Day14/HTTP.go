package main

import (
	"fmt"
	"net/http"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Server is running on port 8080!")
	fmt.Println(r.Method)
}

func HandleHello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello server!")
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", Handle)
	mux.HandleFunc("/hello", HandleHello)

	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}

	server.ListenAndServe()

	fmt.Println("Hello")
}
