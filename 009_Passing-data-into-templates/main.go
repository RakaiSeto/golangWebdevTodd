package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	xi := []int{1, 2, 3, 4, 5}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", xi)
	if err != nil {
		log.Fatalln(err)
	}
}