package product

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type getProductByIDRequest struct {
	ProductID int
}

type getProductsRequest struct {
	Limit  int
	Offset int
}

type getAddProductRequest struct {
	Category     string
	Description  string
	ListPrice    string
	StandardCost string
	ProductCode  string
	ProductName  string
}

type getDeleteProductRequest struct {
	ProductID int
}

type getUpdateProductRequest struct {
	ID           int
	Category     string
	Description  string
	ListPrice    string
	StandardCost string
	ProductCode  string
	ProductName  string
}

type getBestSellersRequest struct{}

// @Summary Lista de Productos
// @Tags Products
// @Accept json
// @Produce json
// @Param id path int true "User Data"
// @Success 200 {object} product.Product "ok"
// @Router /products/{id} [get]
func makeGetProductByIDEndpoint(s Service) endpoint.Endpoint {
	getProductByIDEndpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getProductByIDRequest)
		product, err := s.GetProductByID(&req)

		if err != nil {
			panic(err)
		}

		return product, nil
	}

	return getProductByIDEndpoint
}

// @Summary Producto por ID
// @Tags Products
// @Accept json
// @Produce json
// @Param request body product.getProductsRequest true "User Data"
// @Success 200 {object} product.ProductList "ok"
// @Router /products/paginated [post]
func makeGetProductsEndpoint(s Service) endpoint.Endpoint {
	getProductsEndpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getProductsRequest)
		product, err := s.GetProducts(&req)

		if err != nil {
			panic(err)
		}

		return product, nil
	}

	return getProductsEndpoint
}

// @Summary Insertar un Producto
// @Tags Products
// @Accept json
// @Produce json
// @Param request body product.getAddProductRequest true "User Data"
// @Success 200 {integer} int "ok"
// @Router /products/ [post]
func makeAddProductEndpoint(s Service) endpoint.Endpoint {
	addProductEndpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getAddProductRequest)
		id, err := s.InsertProduct(&req)
		if err != nil {
			panic(err)
		}

		return id, nil
	}

	return addProductEndpoint
}

// @Summary Mejores Ventas Producto
// @Tags Products
// @Accept json
// @Produce json
// @Success 200 {object} product.ProductTopResponse "ok"
// @Router /products/best-sellers [get]
func makeGetBestSellers(s Service) endpoint.Endpoint {
	bestSellersEndpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		bestSellers, err := s.GetBestSellers()
		if err != nil {
			panic(err)
		}

		return bestSellers, nil
	}

	return bestSellersEndpoint
}

// @Summary Actualizar un Producto
// @Tags Products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param request body product.getUpdateProductRequest true "User Data"
// @Success 200 {integer} int "ok"
// @Router /products/{id} [put]
func makeUpdateProductEndpoint(s Service) endpoint.Endpoint {
	updateProductEndpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getUpdateProductRequest)
		id, err := s.UpdateProduct(&req)
		if err != nil {
			panic(err)
		}

		return id, nil
	}

	return updateProductEndpoint
}

// @Summary Eliminar un Productos
// @Tags Products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {integer} int "ok"
// @Router /products/{id} [delete]
func makeDeleteProductEndpoint(s Service) endpoint.Endpoint {
	deleteProductEndpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getDeleteProductRequest)
		rowsAffected, err := s.DeleteProduct(&req)
		if err != nil {
			panic(err)
		}

		return rowsAffected, nil
	}

	return deleteProductEndpoint
}
