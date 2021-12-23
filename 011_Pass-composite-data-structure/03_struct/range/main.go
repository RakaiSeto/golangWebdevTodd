package main

import (
	"text/template"
	"os"
	"log"
)

var tpl *template.Template

type student struct {
	Name string
	FavSub string
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	R28 := student{"Rakai", "PE"}

	err := tpl.Execute(os.Stdout, R28)
	if err != nil {
		log.Fatalln(err)
	}
}
