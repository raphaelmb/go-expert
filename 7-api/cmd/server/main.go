package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"
	"github.com/raphaelmb/go-expert/7-api/configs"
	_ "github.com/raphaelmb/go-expert/7-api/docs"
	"github.com/raphaelmb/go-expert/7-api/internal/entity"
	"github.com/raphaelmb/go-expert/7-api/internal/infra/database"
	"github.com/raphaelmb/go-expert/7-api/internal/infra/webserver/handlers"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// @title           Go Expert API
// @version         1.0
// @description     Product API with authentication.
// @termsOfService  http://swagger.io/terms/

// @contact.name   Raphael Berger
// @contact.url    http://www.github.com/raphaelmb
// @contact.email  raphaelmb@gmail.com

// @license.name 	My License
// @license.url   http://www.github.com/raphaelmb

// @host      localhost:8080
// @BasePath  /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	configs, err := configs.LoadConfig(".")
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
	userDB := database.NewUser(db)
	UserHandler := handlers.NewUserHandler(userDB)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.WithValue("jwt", configs.TokenAuth))
	r.Use(middleware.WithValue("JwtExpiresIn", configs.JwtExpiresIn))

	r.Route("/products", func(r chi.Router) {
		r.Use(jwtauth.Verifier(configs.TokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Post("/", ProductHandler.CreateProduct)
		r.Get("/", ProductHandler.GetAllProducts)
		r.Get("/{id}", ProductHandler.GetProduct)
		r.Put("/{id}", ProductHandler.UpdateProduct)
		r.Delete("/{id}", ProductHandler.DeleteProduct)
	})

	r.Route("/users", func(r chi.Router) {
		r.Post("/", UserHandler.CreateUser)
		r.Post("/generate_token", UserHandler.GetJWT)

	})

	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8080/docs/doc.json")))

	http.ListenAndServe(":8080", r)
}
