package product

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"

	kithttp "github.com/go-kit/kit/transport/http"
)

// MakeHTTPHandler handler for employees
func MakeHTTPHandler(s Service) http.Handler {
	r := chi.NewRouter()

	getProductByIDHandler := kithttp.NewServer(makeGetProductByIDEndpoint(s), getProductByIDRequestDecoder, kithttp.EncodeJSONResponse)
	getProductsHandler := kithttp.NewServer(makeGetProductsEndpoint(s), getProductsRequestDecoder, kithttp.EncodeJSONResponse)
	addProductHandler := kithttp.NewServer(makeAddProductEndpoint(s),
		addProductRequestDecoder, kithttp.EncodeJSONResponse)
	deleteProductHandler := kithttp.NewServer(makeDeleteProductEndpoint(s), deleteProductRequestDecoder, kithttp.EncodeJSONResponse)
	updateProductHandler := kithttp.NewServer(makeUpdateProductEndpoint(s), updateProductRequestDecoder, kithttp.EncodeJSONResponse)
	bestSellersHandler := kithttp.NewServer(makeGetBestSellers(s), getBestSellersRequestDecoder, kithttp.EncodeJSONResponse)
	r.Method(http.MethodGet, "/{id}", getProductByIDHandler)
	r.Method(http.MethodPost, "/paginated", getProductsHandler)
	r.Method(http.MethodPost, "/", addProductHandler)
	r.Method(http.MethodDelete, "/{id}", deleteProductHandler)
	r.Method(http.MethodPut, "/{id}", updateProductHandler)
	r.Method(http.MethodGet, "/best-sellers", bestSellersHandler)

	return r
}

func getProductByIDRequestDecoder(ctx context.Context, r *http.Request) (interface{}, error) {
	productID, _ := strconv.Atoi(chi.URLParam(r, "id"))

	return getProductByIDRequest{
		ProductID: productID,
	}, nil
}

func getProductsRequestDecoder(ctx context.Context, r *http.Request) (interface{}, error) {
	request := getProductsRequest{}

	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		panic(err)
	}

	return getProductsRequest{
		Limit:  request.Limit,
		Offset: request.Offset,
	}, nil
}

func addProductRequestDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	request := getAddProductRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		panic(err)
	}

	return request, nil
}

func deleteProductRequestDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	productID, _ := strconv.Atoi(chi.URLParam(r, "id"))
	return getDeleteProductRequest{
		ProductID: productID,
	}, nil
}

func updateProductRequestDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		panic(err)
	}
	request := getUpdateProductRequest{}
	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		panic(err)
	}

	request.ID = id

	return request, nil
}

func getBestSellersRequestDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	return getBestSellersRequest{}, nil
}
