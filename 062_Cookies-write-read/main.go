package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func set(w http.ResponseWriter, req *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name: "Some-cookie",
		Value: "Anything",
	})
	fmt.Fprintln(w, "COOKIE WRITTEN, CHECK YOUR BROWSER")
	fmt.Fprintln(w, "in chrome go to devs tools")
}

func read(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("Some-cookie")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	fmt.Fprintln(w, "YOUR COOKIE", c)
}