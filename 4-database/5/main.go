package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product
}

type Product struct {
	ID           int `gorm:"primaryKey"`
	Name         string
	Price        float64
	CategoryID   int
	Category     Category
	SerialNumber SerialNumber
	gorm.Model
}

type SerialNumber struct {
	ID        int `gorm:"primaryKey"`
	Number    string
	ProductID int
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})

	// CREATE CATEGORY
	// category := Category{Name: "Eletronics"}
	// db.Create(&category)
	// category := Category{Name: "Kitchen"}
	// db.Create(&category)

	// CREATE PRODUCT
	// product := Product{Name: "Notebook", Price: 1000.00, CategoryID: category.ID}
	// product := Product{Name: "Knife", Price: 100.00, CategoryID: category.ID}
	// db.Create(&product)

	// CREATE SERIAL
	// db.Create(&SerialNumber{Number: "123456", ProductID: product.ID})

	var categories []Category
	err = db.Model(&Category{}).Preload("Products").Preload("Products.SerialNumber").Find(&categories).Error
	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		fmt.Println(category.Name, ":")
		for _, product := range category.Products {
			fmt.Println("-", product.Name, "Serial Number:", product.SerialNumber.Number)
		}
	}
}
