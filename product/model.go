package product

type Product struct {
	ID           int     `json:"id"`
	ProductCode  string  `json:"product_code"`
	ProductName  string  `json:"product_name"`
	Description  string  `json:"description"`
	StandardCost float64 `json:"standard_cost"`
	ListPrice    float64 `json:"list_price"`
	Category     string  `json:"category"`
}

type ProductList struct {
	Data         []*Product `json:"data"`
	TotalRecords int        `json:"total_records"`
}

type ProductTop struct {
	ID          int     `json:"id"`
	ProductName string  `json:"product_name"`
	Sellers     float64 `json:"sellers"`
}

type ProductTopResponse struct {
	Data         []*ProductTop `json:"data"`
	TotalSellers float64       `json:"total_sellers"`
}
