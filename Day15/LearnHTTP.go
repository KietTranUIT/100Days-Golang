package main

import (
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"strings"
)

type User struct {
	Name, Address, Job, Interest string
	Age                          int
}

var user = []User{
	{Name: "Tran Quang Kiet", Address: "Viet Nam", Job: "Student", Interest: "Ping pong, play game", Age: 20},
	{Name: "Tran Nhat Anh Thu", Address: "Viet Nam", Job: "Student", Interest: "Badminton, dance", Age: 10},
	{Name: "John", Address: "Ameriaca", Job: "IT", Interest: "Coding", Age: 25},
	{Name: "Dang Quoc Hung", Address: "Viet Nam", Job: "Student", Interest: "No", Age: 25},
	{Name: "Bui Tan Hung", Address: "Viet Nam", Job: "Student", Interest: "No", Age: 25},
}

func FindUser(query string) []User {
	var result []User
	for _, u := range user {
		if strings.Contains(u.Name, query) {
			result = append(result, u)
		}
	}
	return result
}

func Handle(w http.ResponseWriter, r *http.Request) {
	tmpl.Execute(w, nil)
}

func HandleGetUser(w http.ResponseWriter, r *http.Request) {
	u, err := url.Parse(r.URL.String())

	if err != nil {
		panic(err.Error())
	}

	query := u.Query()
	search := query.Get("q")

	usr := FindUser(search)

	if len(usr) == 0 {
		w.Write([]byte("<h6> No result! </h6>"))
		return
	}

	tmpl.Execute(w, usr)
}

var tmpl = template.Must(template.ParseFiles("index.html"))

func main() {
	port := "8181"
	host := "127.0.0.1"

	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))
	mux.HandleFunc("/", Handle)
	mux.HandleFunc("/search", HandleGetUser)

	server := http.Server{
		Addr:    fmt.Sprintf("%s:%s", host, port),
		Handler: mux,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("Server is running on %s:%s\n", host, port)
}
