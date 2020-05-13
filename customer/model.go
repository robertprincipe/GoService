package customer

type Customer struct {
	ID            int    `json:"id"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	BusinessPhone string `json:"business_phone"`
	Address       string `json:"address"`
	Company       string `json:"company"`
	City          string `json:"city"`
}

type CustomerList struct {
	Data         []*Customer `json:"data"`
	TotalRecords int64       `json:"total_records"`
}
