package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/raphaelmb/go-expert/14-sqlc/internal/db"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	ctx := context.Background()
	conn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	queries := db.New(conn)

	err = queries.CreateCategory(ctx, db.CreateCategoryParams{Name: "Backend", ID: uuid.New().String(), Description: sql.NullString{String: "Backend description"}})
	if err != nil {
		panic(err)
	}

	categories, err := queries.ListCategories(ctx)
	if err != nil {
		panic(err)
	}
	for _, category := range categories {
		fmt.Println(category.ID, category.Name, category.Description.String)
	}

	err = queries.UpdateCategory(ctx, db.UpdateCategoryParams{ID: "", Name: "", Description: sql.NullString{String: ""}})
	if err != nil {
		panic(err)
	}

	err = queries.DeleteCategory(ctx, "")
	if err != nil {
		panic(err)
	}
}
