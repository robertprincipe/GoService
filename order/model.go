package order

type OrderList struct {
	Data         []*OrderItem `json:"data"`
	TotalRecords int64        `json:"total_records"`
}

type OrderItem struct {
	ID         int64              `json:"id"`
	CustomerID int                `json:"customer_id"`
	OrderDate  string             `json:"order_date"`
	StatusID   string             `json:"status_id"`
	StatusName string             `json:"status_name"`
	Customer   string             `json:"customer"`
	Company    string             `json:"company"`
	Address    string             `json:"address"`
	Phone      string             `json:"phone"`
	City       string             `json:"city"`
	Data       []*OrderDetailItem `json:"data"`
}

type OrderDetailItem struct {
	ID          int64   `json:"id"`
	OrderID     int     `json:"order_id"`
	ProductID   int     `json:"product_id"`
	Quantity    float64 `json:"quantity"`
	UnitPrice   float64 `json:"unit_price"`
	ProductName string  `json:"product_name"`
}
