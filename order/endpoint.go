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

type addOrderRequest struct {
	ID           int64
	OrderDate    string
	CustomerID   int
	OrderDetails []addOrderDetailRequest
}

type addOrderDetailRequest struct {
	ID        int64
	OrderID   int64
	ProductID int64
	Quantity  int64
	UnitPrice float64
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
	getOrdersEndpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getOrdersRequest)
		orderList, err := s.GetOrders(&req)
		if err != nil {
			panic(err)
		}
		return orderList, err
	}

	return getOrdersEndpoint
}

func makeAddOrderEdnpoint(s Service) endpoint.Endpoint {
	addOrderEndpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(addOrderRequest)
		orderID, err := s.InsertOrder(&req)
		if err != nil {
			panic(err)
		}
		return orderID, nil
	}

	return addOrderEndpoint
}

func makeUpdateOrderEndpoint(s Service) endpoint.Endpoint {
	updateOrderEnpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(addOrderRequest)
		orderID, err := s.UpdateOrder(&req)
		if err != nil {
			panic(err)
		}
		return orderID, nil
	}
	return updateOrderEnpoint
}
