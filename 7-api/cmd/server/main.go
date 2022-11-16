package main

import (
	"net/http"

	"github.com/raphaelmb/go-expert/7-api/configs"
	"github.com/raphaelmb/go-expert/7-api/internal/entity"
	"github.com/raphaelmb/go-expert/7-api/internal/infra/database"
	"github.com/raphaelmb/go-expert/7-api/internal/infra/webserver/handlers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	_, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(sqlite.Open("teste.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.User{}, &entity.Product{})

	productDB := database.NewProduct(db)
	ProductHandler := handlers.NewProductHandler(productDB)

	http.HandleFunc("/products", ProductHandler.CreateProduct)
	http.ListenAndServe(":8080", nil)
}
