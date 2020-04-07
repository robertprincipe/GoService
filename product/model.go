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
