package order

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type getOrderByIDRequest struct {
	orderID int64
}

type getOrdersRequest struct {
	Limit    int
	Offset   int
	Status   interface{}
	DateFrom interface{}
	DateTo   interface{}
}

func makeGetOrderByIDEndpoint(s Service) endpoint.Endpoint {
	getOrderByIDEndpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getOrderByIDRequest)
		orders, err := s.GetOrderByID(&req)
		if err != nil {
			panic(err)
		}

		return orders, nil
	}

	return getOrderByIDEndpoint
}

func makeGetOrdersEndpoint(s Service) endpoint.Endpoint {
	getOrders := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getOrdersRequest)
		orderList, err := s.GetOrders(&req)
		if err != nil {
			panic(err)
		}
		return orderList, err
	}

	return getOrders
}
