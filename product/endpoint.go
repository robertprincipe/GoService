package product

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type getProductByIDRequest struct {
	ProductID int
}

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
