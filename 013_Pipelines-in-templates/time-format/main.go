package main

import (
	"log"
	"text/template"
	"time"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(ft).ParseFiles("tpl.gohtml"))
}

func dmy(t time.Time) string {
	return t.Format("02-01-2006 03:04PM")
}

var ft = template.FuncMap{
	"fTime" : dmy,
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", time.Now())
	if err != nil {
		log.Fatalln(err)
	}
}