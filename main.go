package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type SignUpData struct{
	FirstName string
	LastName	string
	Email 	string
	Password string
} 

var Data SignUpData

var tpl *template.Template

func main() {
	tpl, _ = tpl.ParseGlob("pages/*.html")
	fmt.Println("http://localhost:8040")

	mux := http.NewServeMux()
	mux.HandleFunc("/", HomeHandler)
	mux.HandleFunc("/signup", SignUp)

	http.ListenAndServe(":8040", mux)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		fmt.Fprint(w, "404 Page Not Found")
		return
	}

	if r.Method != "GET" {
		fmt.Fprint(w, "405 Method Not Allowed")
		return
	}

	
	tpl.ExecuteTemplate(w, "index.html", nil)
	
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		fmt.Fprint(w, "405 Method Not Allowed")
		return
	}

	tmp, err := template.ParseFiles("./pages/signup.html")
	if err != nil {
		fmt.Fprint(w, "500 Server Error")
		return
	}

	err = tmp.Execute(w, nil)
	if err != nil {
		fmt.Fprint(w, http.StatusInternalServerError)
	}
}
