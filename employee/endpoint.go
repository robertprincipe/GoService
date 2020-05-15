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

// @Summary Lista de Empleados
// @Tags Employees
// @Accept json
// @Produce json
// @Param request body employee.getEmployeesRequest true "User Data"
// @Success 200 {object} employee.EmployeeList "ok"
// @Router /employees/paginated [post]
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

// @Summary Empleado por ID
// @Tags Employees
// @Accept json
// @Produce json
// @Param id path int true "Employee ID"
// @Success 200 {object} employee.Employee "ok"
// @Router /employees/{id} [get]
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

// @Summary Mejor Empleado
// @Tags Employees
// @Accept json
// @Produce json
// @Success 200 {object} employee.BestEmployee "ok"
// @Router /employees/best [get]
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

// @Summary Insertar un Empleado
// @Tags Employees
// @Accept json
// @Produce json
// @Param request body employee.getAddEmployeeRequest true "User Data"
// @Success 200 {integer} int "ok"
// @Router /employees/ [post]
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

// @Summary Actualizar un Empleado
// @Tags Employees
// @Accept json
// @Produce json
// @Param id path int true "Employee ID"
// @Param request body employee.updateEmployeeRequest true "User Data"
// @Success 200 {integer} int "ok"
// @Router /employees/{id} [put]
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

// @Summary Eliminar un Empleados
// @Tags Employees
// @Accept json
// @Produce json
// @Param id path int true "Employee ID"
// @Success 200 {integer} int "ok"
// @Router /employees/{id} [delete]
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
