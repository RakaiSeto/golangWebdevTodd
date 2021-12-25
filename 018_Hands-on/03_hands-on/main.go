package main

import (
	"log"
	"os"
	"text/template"
)

type hotel struct {
	Name string
	Address string
	City string
	Zip int
	Region string
}

type hotels []hotel

var tpl *template.Template

func init() {tpl = template.Must(template.ParseFiles("tpl.gohtml"))}

func main() {
	h := hotels{
		hotel{
			"Hotel A", "Ivory street", "Los Angeles", 41356, "Southern",
		},
		hotel{
			"Hotel B", "Coconut boulevard", "Oakland", 46363, "Eastern",
		},
		hotel{
			"Hotel C", "Governor avenue", "Sacramento", 71836, "Southern",
		},
	}

	err := tpl.Execute(os.Stdout, h)
	if err != nil {
		log.Fatalln(err)
	}
}