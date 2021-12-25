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

type menu []food

type restaurant struct {
	RestName string
	Menus menu
}

type restaurants []restaurant

var tpl *template.Template

func init() {tpl = template.Must(template.ParseFiles("tpl.gohtml"))}

func main() {
	r := restaurants{
		restaurant{
			"Warteg Budi",
			menu{
				food{"Gado-gado", "Lunch", 23000},
				food{"Soto Ayam", "Dinner", 26000},
				food{"Bubur Ayam", "Breakfast", 17000},
			},
		},
		restaurant{
			"Feast Steak",
			menu{
				food{"Sirloin", "Lunch", 85000},
				food{"Tenderloin", "Lunch", 75000},
				food{"Chicken Steak", "Lunch", 50000},
			},
		},
		restaurant{
			"McMickey",
			menu{
				food{"Big Mc", "Lunch", 35000},
				food{"Cut o'Fish", "Lunch", 28000},
				food{"Nasi Kuning McM", "Breakfast", 25000},
			},
		},
	}
	err := tpl.Execute(os.Stdout, r)
	if err != nil {
		log.Fatalln(err)
	}
}