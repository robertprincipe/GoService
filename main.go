package main

import (
	"net/http"

	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
	"github.com/robertprincipe/goservice/database"
	"github.com/robertprincipe/goservice/employee"
	"github.com/robertprincipe/goservice/product"
)

func main() {
	db := database.InitConnection()

	defer db.Close()

	var (
		productRepository  = product.NewRepository(db)
		employeeRepository = employee.NewRepository(db)
	)

	var (
		productService  product.Service
		employeeService employee.Service
	)

	productService = product.NewService(productRepository)
	employeeService = employee.NewService(employeeRepository)

	r := chi.NewRouter()

	r.Mount("/products", product.MakeHTTPHandler(productService))
	r.Mount("/employees", employee.MakeHTTPHandler(employeeService))

	http.ListenAndServe(":8008", r)
}
