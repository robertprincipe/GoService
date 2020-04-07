package product

import "database/sql"

type Repository interface {
	GetProductByID(id int) (*Product, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(dbConection *sql.DB) Repository {
	return &repository{db: dbConection}
}

func (r *repository) GetProductByID(id int) (*Product, error) {

	const sql = `-- beginsql
					SELECT id, product_code, product_name, COALESCE(description, ''),
					standard_cost, list_price,
					category FROM products WHERE id = ?
					-- endsql
					`

	row := r.db.QueryRow(sql, id)

	product := &Product{}

	err := row.Scan(&product.ID, &product.ProductCode, &product.ProductName, &product.Description, &product.StandardCost, &product.ListPrice, &product.Category)

	if err != nil {
		panic(err)
	}

	return product, err
}
