package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {io.WriteString(w, "index")}
func dog(w http.ResponseWriter, req *http.Request) {io.WriteString(w, "woof woof")}
func me(w http.ResponseWriter, req *http.Request) {io.WriteString(w, "Hi, i'm Rakai")}