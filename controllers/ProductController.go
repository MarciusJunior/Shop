package controllers

import (
	"Shop/models"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	allProducts := models.FindAllProducts()
	temp.ExecuteTemplate(w, "Index", allProducts)
}

func NewProduct(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "NewProduct", nil)
}

func InsertProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		quantity := r.FormValue("quantity")
		price := r.FormValue("price")
		score := r.FormValue("score")

		convertPrice, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Error converting price", err)
		}

		convertQuantity, err := strconv.Atoi(quantity)
		if err != nil {
			log.Println("Error converting quantity", err)
		}

		convertScore, err := strconv.Atoi(score)
		if err != nil {
			log.Println("Error converting score", err)
		}

		models.CreateNewProduct(name, description, convertQuantity, convertPrice, convertScore)

	}

	http.Redirect(w, r, "/", 301)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	idProduct := r.URL.Query().Get("id")
	models.DeleteProductById(idProduct)
	http.Redirect(w, r, "/", 301)
}

func EditProduct(w http.ResponseWriter, r *http.Request) {
	idProduct := r.URL.Query().Get("id")
	product := models.EditProductById(idProduct)
	temp.ExecuteTemplate(w, "EditProduct", product)
}
