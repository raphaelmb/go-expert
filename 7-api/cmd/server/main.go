package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/products", ProductHandler.CreateProduct)
	r.Get("/products", ProductHandler.GetAllProducts)
	r.Get("/products/{id}", ProductHandler.GetProduct)
	r.Put("/products/{id}", ProductHandler.UpdateProduct)
	r.Delete("/products/{id}", ProductHandler.DeleteProduct)

	http.ListenAndServe(":8080", r)
}
