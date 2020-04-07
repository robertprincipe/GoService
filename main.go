package main

import (
	"net/http"

	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
	"github.com/robertprincipe/goservice/database"
	"github.com/robertprincipe/goservice/product"
)

func main() {
	db := database.InitConnection()

	defer db.Close()

	var productRepository = product.NewRepository(db)
	var productService product.Service
	productService = product.NewService(productRepository)

	r := chi.NewRouter()

	r.Mount("/products", product.MakeHttpHandler(productService))

	http.ListenAndServe(":8008", r)
}
