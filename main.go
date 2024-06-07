package main

import (
	"html/template"
	"net/http"
)

type Product struct {
	Name        string
	Description string
	Quantity    int
	Price       float64
	Score       int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)

}

func index(w http.ResponseWriter, r *http.Request) {
	products := []Product{
		{Name: "Notebook Alienware", Description: "This is the notebook alienware.", Quantity: 10, Price: 2500, Score: 9},
		{"Apple Pencil", "This is the Apple Pencil for iPad.", 27, 190, 8},
		{"Notebook Asus", "16GB de RAM - AMD Ryzen 5 5600X - RTX 4070 TI SUPER", 60, 3990.99, 9},
	}
	temp.ExecuteTemplate(w, "Index", products)
}
