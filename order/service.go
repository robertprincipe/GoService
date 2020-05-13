package order

type Service interface {
	GetOrderByID(param *getOrderByIDRequest) (*OrderItem, error)
	GetOrders(params *getOrdersRequest) (*OrderList, error)
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) GetOrderByID(param *getOrderByIDRequest) (*OrderItem, error) {
	orders, err := s.r.GetOrderByID(param)
	if err != nil {
		panic(err)
	}
	return orders, nil
}

func (s *service) GetOrders(params *getOrdersRequest) (*OrderList, error) {
	orders, err := s.r.GetOrders(params)
	if err != nil {
		panic(err)
	}
	totalRecords, err := s.r.GetTotalOrders(params)
	if err != nil {
		panic(err)
	}
	return &OrderList{
		Data:         orders,
		TotalRecords: totalRecords,
	}, nil
}
