package employee

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

	getEmployeesHandler := kithttp.NewServer(makeGetEmployeesEndpoint(s), getEmployeesRequestDecoder, kithttp.EncodeJSONResponse)
	getEmployeeByIDHandler := kithttp.NewServer(makeGetEmployeeByIDEndpoint(s), getEmployeeByIDRequestDecoder, kithttp.EncodeJSONResponse)
	getBestEmployeeHandler := kithttp.NewServer(makeGetBestEmployeeEndpoint(s), getBestEmployeeRequestDecoder, kithttp.EncodeJSONResponse)
	addEmployeeHandler := kithttp.NewServer(makeAddEmployeeEndpoint(s), addEmployeeRequestDecoder, kithttp.EncodeJSONResponse)
	deleteEmployeeHandler := kithttp.NewServer(makeDeleteEmployeeEndpoint(s), deleteEmployeeRequestDecoder, kithttp.EncodeJSONResponse)
	updateEmployeeHandler := kithttp.NewServer(makeUpdateEmployeeEndpoint(s), updateEmployeeRequestDecoder, kithttp.EncodeJSONResponse)

	r.Method(http.MethodPost, "/paginated", getEmployeesHandler)
	r.Method(http.MethodGet, "/{id}", getEmployeeByIDHandler)
	r.Method(http.MethodGet, "/best", getBestEmployeeHandler)
	r.Method(http.MethodPost, "/", addEmployeeHandler)
	r.Method(http.MethodPut, "/{id}", updateEmployeeHandler)
	r.Method(http.MethodDelete, "/{id}", deleteEmployeeHandler)

	return r
}

func getEmployeesRequestDecoder(ctx context.Context, r *http.Request) (interface{}, error) {
	request := getEmployeesRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		panic(err)
	}
	return request, nil
}

func getEmployeeByIDRequestDecoder(ctx context.Context, r *http.Request) (interface{}, error) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		panic(err)
	}

	return getEmployeeByIDRequest{id}, nil
}

func getBestEmployeeRequestDecoder(ctx context.Context, r *http.Request) (interface{}, error) {
	return getBestEmployeedRequest{}, nil
}

func addEmployeeRequestDecoder(ctx context.Context, r *http.Request) (interface{}, error) {
	request := getAddEmployeeRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		panic(err)
	}
	return request, nil
}

func updateEmployeeRequestDecoder(ctx context.Context, r *http.Request) (interface{}, error) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		panic(err)
	}

	request := updateEmployeeRequest{}
	err = json.NewDecoder(r.Body).Decode(&request)

	return updateEmployeeRequest{
		ID:            id,
		FirstName:     request.FirstName,
		LastName:      request.LastName,
		Company:       request.Company,
		EmailAddress:  request.EmailAddress,
		JobTitle:      request.JobTitle,
		BusinessPhone: request.BusinessPhone,
		HomePhone:     request.HomePhone,
		MobilePhone:   request.MobilePhone,
		FaxNumber:     request.FaxNumber,
		Address:       request.Address,
	}, nil
}

func deleteEmployeeRequestDecoder(ctx context.Context, r *http.Request) (interface{}, error) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		panic(err)
	}

	return deleteEmployeeRequest{id}, nil
}
