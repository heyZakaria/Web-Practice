package main

import (
	"fmt"
	f "fmt"
	"net/http"
)

type Data struct{}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	f.Fprintln(w, "Hello Bro!")
}

func QueryHome(w http.ResponseWriter, r *http.Request) {

	name := r.URL.Query().Get("name")
	if name != "" {
		fmt.Fprint(w, "Hello")
	} else {
		fmt.Fprintf(w, "Hello %s", name)
	}
}

func main() {
	f.Println("http://localhost:8080")



	mux := http.NewServeMux()

	mux.HandleFunc("/set-cookies", )
	mux.HandleFunc("/get-cookies", )

	mux.HandleFunc("/", HomeHandler)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
