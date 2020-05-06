package employee

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type getEmployeesRequest struct {
	Offset int
	Limit  int
}

type getEmployeeByIDRequest struct {
	ID int
}

type getBestEmployeedRequest struct{}

type deleteEmployeeRequest struct {
	EmployeeID int
}

type getAddEmployeeRequest struct {
	FirstName     string
	LastName      string
	Company       string
	EmailAddress  string
	JobTitle      string
	BusinessPhone string
	HomePhone     string
	MobilePhone   string
	FaxNumber     string
	Address       string
}

type updateEmployeeRequest struct {
	ID            int
	FirstName     string
	LastName      string
	Company       string
	EmailAddress  string
	JobTitle      string
	BusinessPhone string
	HomePhone     string
	MobilePhone   string
	FaxNumber     string
	Address       string
}

func makeGetEmployeesEndpoint(s Service) endpoint.Endpoint {
	getEmployeesEndpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getEmployeesRequest)
		data, err := s.GetEmployees(&req)
		if err != nil {
			panic(err)
		}

		return data, err
	}

	return getEmployeesEndpoint
}

func makeGetEmployeeByIDEndpoint(s Service) endpoint.Endpoint {
	getEmployeeByIDEndpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getEmployeeByIDRequest)
		employee, err := s.GetEmployeeByID(&req)
		if err != nil {
			panic(err)
		}

		return employee, nil
	}
	return getEmployeeByIDEndpoint
}

func makeGetBestEmployeeEndpoint(s Service) endpoint.Endpoint {
	getBestEmployee := func(ctx context.Context, request interface{}) (interface{}, error) {
		bestEmployee, err := s.GetBestEmployeed()
		if err != nil {
			panic(err)
		}

		return bestEmployee, nil
	}
	return getBestEmployee
}

func makeAddEmployeeEndpoint(s Service) endpoint.Endpoint {
	addEmployeeEndpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getAddEmployeeRequest)
		rowsAffected, err := s.InsertEmployee(&req)
		if err != nil {
			panic(err)
		}
		return rowsAffected, nil
	}
	return addEmployeeEndpoint
}

func makeDeleteEmployeeEndpoint(s Service) endpoint.Endpoint {
	deleteEmployeeEndpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(deleteEmployeeRequest)
		rowsAffected, err := s.DeleteEmployee(&req)
		if err != nil {
			panic(err)
		}
		return rowsAffected, nil
	}

	return deleteEmployeeEndpoint
}

func makeUpdateEmployeeEndpoint(s Service) endpoint.Endpoint {
	updateEmployeeEndpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(updateEmployeeRequest)
		rowsAffected, err := s.UpdateEmployee(&req)
		if err != nil {
			panic(err)
		}
		return rowsAffected, nil
	}

	return updateEmployeeEndpoint
}
