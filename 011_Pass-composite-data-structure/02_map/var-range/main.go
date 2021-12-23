package main

import (
	"text/template"
	"os"
	"log"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	absen := map[string]int{"Ariq" : 23, "Rili" : 18, "Rakai" : 28,}

	err := tpl.Execute(os.Stdout, absen)
	if err != nil {
		log.Fatalln(err)
	}
}
