package order

type Service interface {
	GetOrderByID(param *getOrderByIDRequest) (*OrderItem, error)
	GetOrders(params *getOrdersRequest) (*OrderList, error)
	InsertOrder(params *addOrderRequest) (int64, error)
	UpdateOrder(params *addOrderRequest) (int64, error)
	DeleteOrderDetail(param *deleteOrderDetailRequest) (int64, error)
	DeleteOrder(param *deleteOrderRequest) (int64, error)
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

func (s *service) InsertOrder(params *addOrderRequest) (int64, error) {
	id, err := s.r.InsertOrder(params)
	if err != nil {
		panic(err)
	}

	for _, detail := range params.OrderDetails {
		detail.OrderID = id
		_, err := s.r.InsertOrderDetail(&detail)
		if err != nil {
			panic(err)
		}
	}
	return id, nil
}

func (s *service) UpdateOrder(params *addOrderRequest) (int64, error) {
	orderId, err := s.r.UpdateOrder(params)
	if err != nil {
		panic(err)
	}

	for _, detail := range params.OrderDetails {
		detail.OrderID = orderId
		if detail.ID == 0 {
			_, err := s.r.InsertOrderDetail(&detail)
			if err != nil {
				panic(err)
			}
		} else {
			_, err := s.r.UpdateOrderDetail(&detail)
			if err != nil {
				panic(err)
			}
		}
	}

	return orderId, nil
}

func (s *service) DeleteOrderDetail(param *deleteOrderDetailRequest) (int64, error) {
	rowsAffected, err := s.r.DeleteOrderDetail(param)
	if err != nil {
		panic(err)
	}
	return rowsAffected, nil
}

func (s *service) DeleteOrder(param *deleteOrderRequest) (int64, error) {
	_, err := s.r.DeleteOrderDetailByOrderID(param)
	if err != nil {
		panic(err)
	}
	return s.r.DeleteOrder(param)
}
