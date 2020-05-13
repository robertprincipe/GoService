package customer

import "database/sql"

type Repository interface {
	GetCustomers(params *getCustomersRequest) ([]*Customer, error)
	GetTotalCustomer() (int64, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db}
}

func (r *repository) GetCustomers(params *getCustomersRequest) ([]*Customer, error) {
	const query = `
	-- beginsql
	SELECT id, first_name, last_name, business_phone, address, company, city FROM customers
	 	LIMIT ? OFFSET ?
	-- endsql
	`
	results, err := r.db.Query(query, params.Limit, params.Offset)
	if err != nil {
		panic(err)
	}
	var customers []*Customer
	for results.Next() {
		customer := Customer{}
		err = results.Scan(&customer.ID, &customer.FirstName, &customer.LastName, &customer.BusinessPhone,
			&customer.Address, &customer.Company, &customer.City)
		if err != nil {
			panic(err)
		}
		customers = append(customers, &customer)
	}

	return customers, nil
}

func (r *repository) GetTotalCustomer() (int64, error) {
	const query = `
	-- beginsql
	SELECT COUNT(id) FROM customers
	-- endsql
	`
	result := r.db.QueryRow(query)
	var totalRecords int64
	err := result.Scan(&totalRecords)
	if err != nil {
		panic(err)
	}

	return totalRecords, nil
}
