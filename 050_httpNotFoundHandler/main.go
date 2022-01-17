package main

import (
	"net/http"
	"fmt"
)

func main() {
	http.HandleFunc("/", dog)
	http.Handle("/nothing", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request){
	fmt.Println(req.URL)
	fmt.Fprintln(w, "look at the terminal")
}