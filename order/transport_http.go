package order

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	kithttp "github.com/go-kit/kit/transport/http"
)

func MakeHTTPHandler(s Service) http.Handler {
	r := chi.NewRouter()

	getOrderByIDHandler := kithttp.NewServer(makeGetOrderByIDEndpoint(s), getOrderByIDRequestDecoder, kithttp.EncodeJSONResponse)
	getOrdersHandler := kithttp.NewServer(makeGetOrdersEndpoint(s), getOrdersRequestDecoder, kithttp.EncodeJSONResponse)
	addOrderHandler := kithttp.NewServer(makeAddOrderEdnpoint(s), addOrderRequestDecoder, kithttp.EncodeJSONResponse)
	updateOrderHandler := kithttp.NewServer(makeUpdateOrderEndpoint(s), updateOrderRequestDecoder, kithttp.EncodeJSONResponse)
	deleteOrderDetailHandler := kithttp.NewServer(makeDeleteOrderDetailEndpoint(s), deleteOrderDetailRequestDecoder, kithttp.EncodeJSONResponse)
	deleteOrderHandler := kithttp.NewServer(makeDeleteOrderEndpoint(s), deleteOrderRequestDecoder, kithttp.EncodeJSONResponse)
	r.Method(http.MethodGet, "/{id}", getOrderByIDHandler)
	r.Method(http.MethodPost, "/paginated", getOrdersHandler)
	r.Method(http.MethodPost, "/", addOrderHandler)
	r.Method(http.MethodPut, "/", updateOrderHandler)
	r.Method(http.MethodDelete, "/{orderID}/order-detail/{orderDetailID}", deleteOrderDetailHandler)
	r.Method(http.MethodDelete, "/{orderID}", deleteOrderHandler)
	return r
}

func getOrderByIDRequestDecoder(ctx context.Context, r *http.Request) (interface{}, error) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		panic(err)
	}
	return getOrderByIDRequest{id}, nil
}

func getOrdersRequestDecoder(ctx context.Context, r *http.Request) (interface{}, error) {
	request := getOrdersRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		panic(err)
	}
	return request, nil
}

func addOrderRequestDecoder(ctx context.Context, r *http.Request) (interface{}, error) {
	request := addOrderRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		panic(err)
	}
	return request, nil
}

func updateOrderRequestDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	request := addOrderRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		panic(err)
	}
	return request, nil
}

func deleteOrderDetailRequestDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	return deleteOrderDetailRequest{
		OrderDetailID: chi.URLParam(r, "orderDetailID"),
	}, nil
}

func deleteOrderRequestDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	return deleteOrderRequest{
		OrderID: chi.URLParam(r, "orderID"),
	}, nil
}
