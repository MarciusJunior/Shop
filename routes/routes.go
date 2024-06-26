package routes

import (
	"Shop/controllers"
	"net/http"
)

func LoadRoutes() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/newProduct", controllers.NewProduct)
	http.HandleFunc("/insertProduct", controllers.InsertProduct)
	http.HandleFunc("/delete", controllers.DeleteProduct)
	http.HandleFunc("/editProduct", controllers.EditProduct)
	http.HandleFunc("/update", controllers.UpdateProduct)
}
