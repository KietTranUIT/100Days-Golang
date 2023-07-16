package main

import (
	"fmt"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Welcome to HTTP server")
}

const server_ip = "127.0.0.1"
const server_port = 8080

func main() {

	mux := http.NewServeMux()

	server := http.Server {
		Addr: fmt.Sprintf("%s:%d", server_ip, server_port),
		Handler: mux,
	}

	mux.HandleFunc("/", Handler)

	err := server.ListenAndServe()

	fmt.Println("Server staring on port 8080")


	if err != nil {
		panic(err.Error())
	}

}