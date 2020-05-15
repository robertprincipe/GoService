package customer

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/segmentio/encoding/json"
)

func MakeHTTPHandler(s Service) http.Handler {
	r := chi.NewRouter()

	getCustomersHandler := kithttp.NewServer(makeGetCustomersEnpoint(s), getCustomersRequestDecoder, kithttp.EncodeJSONResponse)

	r.Method(http.MethodPost, "/paginated", getCustomersHandler)

	return r
}

func getCustomersRequestDecoder(ctx context.Context, r *http.Request) (interface{}, error) {
	request := getCustomersRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		panic(err)
	}

	return request, nil
}
