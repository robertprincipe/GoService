package product

type Service interface {
	GetProductByID(param *getProductByIDRequest) (*Product, error)
	GetProducts(params *getProductsRequest) (*ProductList, error)
	InsertProduct(params *getAddProductRequest) (int64, error)
	UpdateProduct(params *getUpdateProductRequest) (int64, error)
	DeleteProduct(param *getDeleteProductRequest) (int64, error)
	GetBestSellers() (*ProductTopResponse, error)
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

func (s *service) GetProducts(params *getProductsRequest) (*ProductList, error) {
	products, err := s.r.GetProducts(params)
	if err != nil {
		panic(err)
	}

	totalProducts, err := s.r.GetTotalProducts()

	if err != nil {
		panic(err)
	}

	return &ProductList{Data: products, TotalRecords: totalProducts}, nil
}

func (s *service) InsertProduct(params *getAddProductRequest) (int64, error) {

	return s.r.InsertProduct(params)
}

func (s *service) DeleteProduct(param *getDeleteProductRequest) (int64, error) {
	return s.r.DeleteProduct(param.ProductID)
}

func (s *service) UpdateProduct(params *getUpdateProductRequest) (int64, error) {
	return s.r.UpdateProduct(params)
}

func (s *service) GetBestSellers() (*ProductTopResponse, error) {
	bestSellers, err := s.r.GetBestSellers()
	if err != nil {
		panic(err)
	}
	totalSellers, err := s.r.GetTotalSellers()

	if err != nil {
		panic(err)
	}

	return &ProductTopResponse{
		Data:         bestSellers,
		TotalSellers: totalSellers,
	}, nil

}
