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
