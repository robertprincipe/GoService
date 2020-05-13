package main

import (
	"net/http"

	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
	"github.com/robertprincipe/goservice/customer"
	"github.com/robertprincipe/goservice/database"
	"github.com/robertprincipe/goservice/employee"
	"github.com/robertprincipe/goservice/order"
	"github.com/robertprincipe/goservice/product"
)

func main() {
	db := database.InitConnection()

	defer db.Close()

	var (
		productRepository  = product.NewRepository(db)
		employeeRepository = employee.NewRepository(db)
		customerRepository = customer.NewRepository(db)
		orderRepository    = order.NewRepository(db)
	)

	var (
		productService  product.Service
		employeeService employee.Service
		customerService customer.Service
		orderService    order.Service
	)

	productService = product.NewService(productRepository)
	employeeService = employee.NewService(employeeRepository)
	customerService = customer.NewService(customerRepository)
	orderService = order.NewService(orderRepository)

	r := chi.NewRouter()

	r.Mount("/products", product.MakeHTTPHandler(productService))
	r.Mount("/employees", employee.MakeHTTPHandler(employeeService))
	r.Mount("/customers", customer.MakeHTTPHandler(customerService))
	r.Mount("/orders", order.MakeHTTPHandler(orderService))

	http.ListenAndServe(":8008", r)
}
