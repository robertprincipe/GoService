package employee

import "database/sql"

type Repository interface {
	GetEmployees(params *getEmployeesRequest) ([]*Employee, error)
	GetTotalEmployees() (int64, error)
	GetEmployeeByID(param *getEmployeeByIDRequest) (*Employee, error)
	GetBestEmployee() (*BestEmployee, error)
	InsertEmployee(params *getAddEmployeeRequest) (int64, error)
	DeleteEmployee(param *deleteEmployeeRequest) (int64, error)
	UpdateEmployee(params *updateEmployeeRequest) (int64, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db}
}

func (r *repository) GetEmployees(params *getEmployeesRequest) ([]*Employee, error) {
	const query = `
	-- beginsql
	SELECT id, first_name, last_name, company, job_title, email_address, business_phone, home_phone, COALESCE(mobile_phone , ''),
	fax_number, address FROM employees LIMIT ? OFFSET ?
	-- endsql
	`

	results, err := r.db.Query(query, params.Limit, params.Offset)

	if err != nil {
		panic(err)
	}

	var employees []*Employee

	for results.Next() {
		employee := Employee{}
		err := results.Scan(&employee.ID, &employee.FirstName, &employee.LastName, &employee.Company, &employee.JobTitle,
			&employee.EmailAddress, &employee.BusinessPhone, &employee.HomePhone, &employee.MobilePhone,
			&employee.FaxNumber, &employee.Address)
		if err != nil {
			panic(err)
		}

		employees = append(employees, &employee)
	}

	return employees, nil
}

func (r *repository) GetTotalEmployees() (int64, error) {
	const query = `
	-- beginsql
	SELECT COUNT(*) FROM employees
	-- endsql
	`
	result := r.db.QueryRow(query)
	var totalEmployees int64
	err := result.Scan(&totalEmployees)

	if err != nil {
		panic(err)
	}

	return totalEmployees, nil
}

func (r *repository) GetEmployeeByID(param *getEmployeeByIDRequest) (*Employee, error) {
	const query = `
	-- beginsql
	SELECT id, first_name, last_name, company, job_title, email_address, business_phone, home_phone, COALESCE(mobile_phone , ''),
	fax_number, address FROM employees WHERE id = ?
	-- endsql
	`
	result := r.db.QueryRow(query, param.ID)
	employee := &Employee{}
	err := result.Scan(&employee.ID, &employee.FirstName, &employee.LastName, &employee.Company, &employee.JobTitle,
		&employee.EmailAddress, &employee.BusinessPhone, &employee.HomePhone, &employee.MobilePhone,
		&employee.FaxNumber, &employee.Address)
	if err != nil {
		panic(err)
	}
	return employee, err
}

func (r *repository) GetBestEmployee() (*BestEmployee, error) {
	const query = `
	-- beginsql
	SELECT 
		e.id,
		e.first_name, e.last_name,
		COUNT(e.id) AS total_sellers
	FROM orders o
	INNER JOIN employees e ON o.employee_id = e.id
	GROUP BY o.employee_id
	ORDER BY total_sellers DESC
	LIMIT 1
	-- endsql
	`
	result := r.db.QueryRow(query)
	bestEmployee := &BestEmployee{}
	err := result.Scan(&bestEmployee.ID, &bestEmployee.FirstName, &bestEmployee.LastName, &bestEmployee.TotalSellers)
	if err != nil {
		panic(err)
	}

	return bestEmployee, nil
}

func (r *repository) InsertEmployee(params *getAddEmployeeRequest) (int64, error) {
	const query = `
	-- beginsql
	INSERT INTO employees
	(first_name, last_name, company, job_title, email_address, business_phone, home_phone, mobile_phone,
	fax_number, address) 
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	-- endsql
	`
	result, err := r.db.Exec(query, &params.FirstName, &params.LastName, &params.Company, &params.JobTitle,
		&params.EmailAddress, &params.BusinessPhone, &params.HomePhone, &params.MobilePhone,
		&params.FaxNumber, &params.Address)
	if err != nil {
		panic(err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}
	return rowsAffected, nil
}

func (r *repository) UpdateEmployee(params *updateEmployeeRequest) (int64, error) {
	const query = `
	-- beginsql
	UPDATE employees SET first_name = ?, last_name = ?, company = ?, job_title = ?, email_address = ?, 
	business_phone = ?, home_phone = ?, mobile_phone = ?,
	fax_number = ?, address = ? WHERE id = ?
	-- endsql
	`
	result, err := r.db.Exec(query, &params.FirstName, &params.LastName, &params.Company, &params.JobTitle, &params.EmailAddress,
		&params.BusinessPhone, &params.HomePhone, &params.MobilePhone, &params.FaxNumber, &params.Address, &params.ID)
	if err != nil {
		panic(err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}
	return rowsAffected, nil
}

func (r *repository) DeleteEmployee(param *deleteEmployeeRequest) (int64, error) {
	const query = `
	-- beginsql
	DELETE FROM employees WHERE id = ?
	-- endsql
	`
	result, err := r.db.Exec(query, param.EmployeeID)
	if err != nil {
		panic(err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}
	return rowsAffected, nil
}
