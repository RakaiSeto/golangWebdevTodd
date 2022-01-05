package main

import (
	"io"
	"log"
	"net/http"
	"text/template"
)

func main() {
	http.Handle("/", http.HandlerFunc(index))
	http.Handle("/dog/", http.HandlerFunc(doggy))
	http.Handle("/tm/", http.HandlerFunc(rakai))
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {io.WriteString(w, "This is the home page")}
func doggy(w http.ResponseWriter, req *http.Request) {io.WriteString(w, "Who let the dogs out????")}
func rakai(w http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("something.gohtml")
	if err != nil {
		log.Fatalln("Parse won't go")
	}

	err = tpl.ExecuteTemplate(w, "something.gohtml", "Rakai Seto")
}