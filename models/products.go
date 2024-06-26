package models

import (
	"Shop/db"
)

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

		p.Id = id
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

func DeleteProductById(id string) {
	db := db.ConnectDatabase()

	deleteProduct, err := db.Prepare("delete from products where id=$1")

	if err != nil {
		panic(err.Error())
	}

	deleteProduct.Exec(id)

	defer db.Close()
}

func EditProductById(id string) Product {
	db := db.ConnectDatabase()

	productFromData, err := db.Query("SELECT * FROM products where id=$1", id)
	if err != nil {
		panic(err.Error())
	}

	productForEdit := Product{}

	for productFromData.Next() {
		var id, quantity, score int
		var name, description string
		var price float64

		err = productFromData.Scan(&id, &name, &description, &quantity, &price, &score)
		if err != nil {
			panic(err.Error())
		}

		productForEdit.Id = id
		productForEdit.Name = name
		productForEdit.Description = description
		productForEdit.Quantity = quantity
		productForEdit.Price = price
		productForEdit.Score = score
	}

	defer db.Close()
	return productForEdit
}

func UpdateProduct(id int, name, description string, quantity int, price float64, score int) {
	db := db.ConnectDatabase()

	UpdateProduct, err := db.Prepare("update products set name=$1, description=$2, quantity=$3, price=$4, score=$5 where id=$6")
	if err != nil {
		panic(err.Error())
	}
	UpdateProduct.Exec(name, description, quantity, price, score, id)
	defer db.Close()
}
