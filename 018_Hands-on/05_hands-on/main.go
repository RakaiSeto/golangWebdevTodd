package main

import (
	"log"
	"os"
	"text/template"
)

type food struct{
	Name, Meal string
	Price int
}

type foods []food

var tpl *template.Template

func init() {tpl = template.Must(template.ParseFiles("tpl.gohtml"))}

func main() {
	f := foods{
		food{
			"Bubur ayam", "Breakfast", 20000, 
		},
		food{
			"Bakso sapi", "Lunch", 25000, 
		},
		food{
			"Gado-gado", "Lunch", 25000, 
		},
		food{
			"Soto ayam", "Breakfast", 30000, 
		},
	}

	err := tpl.Execute(os.Stdout, f)
	if err != nil {
		log.Fatalln(err)
	}
}