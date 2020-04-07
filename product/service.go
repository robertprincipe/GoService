package product

type Service interface {
	GetProductByID(param *getProductByIDRequest) (*Product, error)
}

type service struct {
	r Repository
}

func NewService(repo Repository) Service {
	return &service{r: repo}
}

func (s *service) GetProductByID(param *getProductByIDRequest) (*Product, error) {
	return s.r.GetProductByID(param.ProductID)
}
