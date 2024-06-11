package main

import (
	"Shop/routes"
	"net/http"
)

func main() {
	routes.LoadRoutes()
	http.ListenAndServe(":8080", nil)
}
