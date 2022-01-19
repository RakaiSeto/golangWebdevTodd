package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/set", set)
	http.HandleFunc("/read", read)
	http.HandleFunc("/expire", expire)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, `<h1><a href="/set">set cookie</a></h1>`)
}

func set(w http.ResponseWriter, req *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name: "Cookie",
		Value: "Something",
	})
	fmt.Fprintln(w, `<h1><a href="/read">read cookie</a></h1>`)
}

func read(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("Cookie")
	if err != nil {
		http.Redirect(w, req, "/set", http.StatusSeeOther)
		return
	}
	fmt.Fprintf(w, `<h1>Your cookie<br>%v</h1><br><h1><a href="/expire">expire</a></h1>`, c)
}

func expire(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("Cookie")
	if err != nil {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	c.MaxAge = -1
	http.SetCookie(w, c)
	http.Redirect(w, req, "/", http.StatusSeeOther)
}