package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"io"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {

	var s string
	fmt.Println(req.Method)
	if req.Method == http.MethodPost {

		f, h, err := req.FormFile("q")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			fmt.Println(h)
		}

		bs, err := ioutil.ReadAll(f)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		s = string(bs)
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
	<form method="POST" enctype="multipart/form-data">
	<input type="file" name="q">
	<input type="submit">
	</form>
	<br>`+s)

}