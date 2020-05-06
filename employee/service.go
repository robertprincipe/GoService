package employee

type Service interface {
	GetEmployees(params *getEmployeesRequest) (*EmployeeList, error)
	GetEmployeeByID(param *getEmployeeByIDRequest) (*Employee, error)
	GetBestEmployeed() (*BestEmployee, error)
	InsertEmployee(params *getAddEmployeeRequest) (int64, error)
	UpdateEmployee(params *updateEmployeeRequest) (int64, error)
	DeleteEmployee(param *deleteEmployeeRequest) (int64, error)
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s service) GetEmployees(params *getEmployeesRequest) (*EmployeeList, error) {
	employees, err := s.r.GetEmployees(params)

	if err != nil {
		panic(err)
	}

	totalEmployees, err := s.r.GetTotalEmployees()
	if err != nil {
		panic(err)
	}

	return &EmployeeList{
		Data:         employees,
		TotalRecords: totalEmployees,
	}, nil
}

func (s service) GetEmployeeByID(param *getEmployeeByIDRequest) (*Employee, error) {
	employee, err := s.r.GetEmployeeByID(param)
	if err != nil {
		panic(err)
	}

	return employee, nil
}

func (s service) GetBestEmployeed() (*BestEmployee, error) {
	bestEmployee, err := s.r.GetBestEmployee()
	if err != nil {
		panic(err)
	}
	return bestEmployee, nil
}

func (s service) InsertEmployee(params *getAddEmployeeRequest) (int64, error) {
	rowsAffected, err := s.r.InsertEmployee(params)
	if err != nil {
		panic(err)
	}
	return rowsAffected, nil
}

func (s service) DeleteEmployee(param *deleteEmployeeRequest) (int64, error) {
	rowsAffected, err := s.r.DeleteEmployee(param)
	if err != nil {
		panic(err)
	}

	return rowsAffected, nil
}

func (s service) UpdateEmployee(params *updateEmployeeRequest) (int64, error) {
	rowsAffected, err := s.r.UpdateEmployee(params)
	if err != nil {
		panic(err)
	}

	return rowsAffected, nil
}
