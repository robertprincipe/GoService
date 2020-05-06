package employee

type Employee struct {
	ID            int    `json:"id"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	Company       string `json:"company"`
	EmailAddress  string `json:"email_address"`
	JobTitle      string `json:"job_title"`
	BusinessPhone string `json:"business_phone"`
	HomePhone     string `json:"home_phone"`
	MobilePhone   string `json:"mobile_phone"`
	FaxNumber     string `json:"fax_number"`
	Address       string `json:"address"`
}

type EmployeeList struct {
	Data         []*Employee `json:"data"`
	TotalRecords int64       `json:"total_records"`
}

type BestEmployee struct {
	ID           int    `json:"id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	TotalSellers int    `json:"total_sellers"`
}
