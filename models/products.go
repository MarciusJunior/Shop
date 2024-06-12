package models

import "Shop/db"

type Product struct {
	Id          int
	Name        string
	Description string
	Quantity    int
	Price       float64
	Score       int
}

func FindAllProducts() []Product {
	dbase := db.ConnectDatabase()

	selectAllProducts, err := dbase.Query("select * from products")
	if err != nil {
		panic(err.Error())
	}

	p := Product{}
	products := []Product{}

	for selectAllProducts.Next() {
		var id, quantity, score int
		var name, description string
		var price float64

		err = selectAllProducts.Scan(&id, &name, &description, &quantity, &price, &score)
		if err != nil {
			panic(err.Error())
		}

		p.Name = name
		p.Description = description
		p.Quantity = quantity
		p.Price = price
		p.Score = score

		products = append(products, p)
	}
	defer dbase.Close()
	return products
}

func CreateNewProduct(name, description string, quantity int, price float64, score int) {
	db := db.ConnectDatabase()

	insertData, err := db.Prepare("insert into products(name, description, quantity, price, score) values($1, $2, $3, $4, $5)")
	if err != nil {
		panic(err.Error())
	}

	insertData.Exec(name, description, quantity, price, score)

	defer db.Close()
}
