package main

import (
	"fmt"
	"net/http"
	"github.com/satori/go.uuid"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
		cookie, err := req.Cookie("session-id")
		if err != nil {
			id, err := uuid.NewV4()
			if err != nil {
				fmt.Println(err)
			}
			cookie = &http.Cookie{
				Name: "session-id",
				Value: id.String(),
				HttpOnly: true,
			}
			http.SetCookie(w, cookie)
		}
		fmt.Println(cookie)
}