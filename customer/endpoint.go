package customer

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type getCustomersRequest struct {
	Offset int
	Limit  int
}

// @Summary Lista de Clientes
// @Tags Customers
// @Accept json
// @Produce json
// @Param request body customer.getCustomersRequest true "User Data"
// @Success 200 {object} customer.CustomerList "ok"
// @Router /customers/paginated [post]
func makeGetCustomersEnpoint(s Service) endpoint.Endpoint {
	getCustomersEnpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getCustomersRequest)
		data, err := s.GetCustomers(&req)
		if err != nil {
			panic(err)
		}

		return data, nil
	}

	return getCustomersEnpoint
}
