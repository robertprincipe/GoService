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
	r.Method(http.MethodGet, "/{id}", getOrderByIDHandler)
	r.Method(http.MethodPost, "/paginated", getOrdersHandler)
	r.Method(http.MethodPost, "/", addOrderHandler)
	r.Method(http.MethodPut, "/", updateOrderHandler)
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
