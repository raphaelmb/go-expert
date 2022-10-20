package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Price float64
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// CREATE
	db.Create(&Product{Name: "Notebook", Price: 1000.00})

	// CREATE BATCH
	// products := []Product{
	// 	{Name: "Notebook", Price: 1000.00},
	// 	{Name: "Mouse", Price: 50.0},
	// 	{Name: "Keyboard", Price: 100.00},
	// }
	// db.Create(&products)

	// SELECT ONE
	// var product Product
	// db.First(&product, 1)
	// db.First(&product, "name = ?", "Mouse")
	// fmt.Println(product)

	// SELECT ALL
	// var products []Product
	// db.Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }
}
