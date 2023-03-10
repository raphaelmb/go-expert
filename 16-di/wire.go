//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"
	"github.com/google/wire"
	"github.com/raphaelmb/go-expert/16-di/product"
)

var setRepositoryDependency = wire.NewSet(
	product.NewProductRepository,
	wire.Bind(new(product.ProductRepositoryInterface), new(*product.ProductRepository)),
)

func NewUseCase(db *sql.DB) *product.ProductUseCase {
	wire.Build(
		setRepositoryDependency,
		product.NewProductUseCase,
	)
	return &product.ProductUseCase{}
}
