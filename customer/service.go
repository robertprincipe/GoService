package customer

type Service interface {
	GetCustomers(params *getCustomersRequest) (*CustomerList, error)
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s service) GetCustomers(params *getCustomersRequest) (*CustomerList, error) {
	customers, err := s.r.GetCustomers(params)
	if err != nil {
		panic(err)
	}
	totalRecords, err := s.r.GetTotalCustomer()
	if err != nil {
		panic(err)
	}
	return &CustomerList{
		Data:         customers,
		TotalRecords: totalRecords,
	}, nil
}
