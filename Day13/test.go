package main

import (
	"fmt"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello everyone to Website of me!")
}

func main() {
	port := "8080"
	ip := "127.0.0.1"

	mux := http.NewServeMux()

	mux.HandleFunc("/", Handler)

	server := http.Server{
		Addr:    ip + ":" + port,
		Handler: mux,
	}

	fmt.Println("Server is running on port 8080!")

	err := server.ListenAndServe()
	if err != nil {
		panic(err.Error())
	}

}
