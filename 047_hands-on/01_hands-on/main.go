package main

import (
	"io"
	"log"
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/cat", cat)
	http.HandleFunc("/cat.jpg", kucing)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request){
	io.WriteString(w, "foo ran")
}

func cat(w http.ResponseWriter, req *http.Request){
	tpl, err := template.ParseFiles("cat.gohtml")
	if err != nil {
		log.Fatal(err)
	}

	tpl.ExecuteTemplate(w, "cat.gohtml", nil)
}

func kucing(w http.ResponseWriter, req *http.Request){
	http.ServeFile(w, req, "cat.jpg")
}