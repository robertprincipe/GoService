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

type deleteOrderDetailRequest struct {
	OrderDetailID string
}

type deleteOrderRequest struct {
	OrderID string
}

// @Summary Orden por ID
// @Tags Orders
// @Accept json
// @Produce json
// @Param id path int true "Order ID"
// @Success 200 {object} order.OrderItem "ok"
// @Router /orders/{id} [get]
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

// @Summary Lista de Ordenes
// @Tags Orders
// @Accept json
// @Produce json
// @Param request body order.getOrdersRequest true "User Data"
// @Success 200 {object} order.OrderList "ok"
// @Router /orders/paginated [post]
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

// @Summary Insertar Orden
// @Tags Orders
// @Accept json
// @Produce json
// @Param request body order.addOrderRequest true "User Data"
// @Success 200 {integer} int "ok"
// @Router /orders/ [post]
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

// @Summary Actulizar una Orden
// @Tags Orders
// @Accept json
// @Produce json
// @Param request body order.addOrderRequest true "User Data"
// @Success 200 {interger} int "ok"
// @Router /orders/ [put]
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

// @Summary Eliminar detalle de Ordenes
// @Tags Orders
// @Accept json
// @Produce json
// @Param orderDetailID path int true "ID detalle orden"
// @Success 200 {int} int "ok"
// @Router /orders/{orderID}/order-detail/{orderDetailID} [delete]
func makeDeleteOrderDetailEndpoint(s Service) endpoint.Endpoint {
	deleteOrderDetailEnpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(deleteOrderDetailRequest)
		rowsAffected, err := s.DeleteOrderDetail(&req)
		if err != nil {
			panic(err)
		}
		return rowsAffected, nil
	}
	return deleteOrderDetailEnpoint
}

// @Summary Eliminar una Orden
// @Tags Orders
// @Accept json
// @Produce json
// @Param orderID path int true "Orden ID"
// @Success 200 {integer} int "ok"
// @Router /orders/{orderID} [delete]
func makeDeleteOrderEndpoint(s Service) endpoint.Endpoint {
	deleteOrderEnpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(deleteOrderRequest)
		rowsAffected, err := s.DeleteOrder(&req)
		if err != nil {
			panic(err)
		}
		return rowsAffected, nil
	}
	return deleteOrderEnpoint
}
