package product

import "database/sql"

type Repository interface {
	GetProductByID(id int) (*Product, error)
	GetProducts(params *getProductsRequest) ([]*Product, error)
	GetTotalProducts() (int, error)
	InsertProduct(params *getAddProductRequest) (int64, error)
	UpdateProduct(params *getUpdateProductRequest) (int64, error)
	DeleteProduct(id int) (int64, error)
	GetBestSellers() ([]*ProductTop, error)
	GetTotalSellers() (float64, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(dbConection *sql.DB) Repository {
	return &repository{db: dbConection}
}

func (r *repository) GetProductByID(id int) (*Product, error) {

	const query = `-- beginsql
					SELECT id, product_code, product_name, COALESCE(description, ''),
					standard_cost, list_price,
					category FROM products WHERE id = ?
					-- endsql
					`

	row := r.db.QueryRow(query, id)

	product := &Product{}

	err := row.Scan(&product.ID, &product.ProductCode, &product.ProductName, &product.Description, &product.StandardCost, &product.ListPrice, &product.Category)

	if err != nil {
		panic(err)
	}

	return product, err
}

func (r *repository) GetProducts(params *getProductsRequest) ([]*Product, error) {
	const query = `
		-- beginsql
		SELECT id, product_code, product_name, COALESCE(description, ''), 
		standard_cost, list_price, category 
		FROM products ORDER BY id LIMIT ? OFFSET ?
		-- endsql
	`

	results, err := r.db.Query(query, params.Limit, params.Offset)
	if err != nil {
		panic(err)
	}
	var products []*Product
	for results.Next() {
		product := &Product{}
		err = results.Scan(&product.ID, &product.ProductCode, &product.ProductName, &product.Description, &product.StandardCost,
			&product.ListPrice, &product.Category)
		if err != nil {
			panic(err)
		}

		products = append(products, product)
	}

	return products, nil
}

func (r *repository) GetTotalProducts() (int, error) {
	const query = `
		-- beginsql
		SELECT COUNT(*)	FROM products
		-- endsql
	`
	var total int
	row := r.db.QueryRow(query)
	err := row.Scan(&total)
	if err != nil {
		panic(err)
	}
	return total, nil
}

func (r *repository) InsertProduct(params *getAddProductRequest) (int64, error) {
	const query = `
	-- beginsql
	INSERT INTO products
	(product_code,product_name,category,description,list_price,standard_cost)
	VALUES(?,?,?,?,?,?)
	-- endsql
	`

	result, err := r.db.Exec(query, params.ProductCode, params.ProductName,
		params.Category, params.Description, params.ListPrice, params.StandardCost)

	if err != nil {
		panic(err)
	}

	productId, _ := result.LastInsertId()

	return productId, nil
}

func (r *repository) UpdateProduct(params *getUpdateProductRequest) (int64, error) {
	const query = `
	-- beginsql
	UPDATE products SET product_code = ?, product_name = ?, description = ?, standard_cost = ?, list_price = ?, category = ? WHERE id = ?
	-- endsql
	`
	result, err := r.db.Exec(query, params.ProductCode, params.ProductName, params.Description,
		params.StandardCost, params.ListPrice, params.Category, params.ID)
	if err != nil {
		panic(err)
	}

	id, err := result.LastInsertId()

	if err != nil {
		panic(err)
	}

	return id, nil
}

func (r *repository) DeleteProduct(id int) (int64, error) {
	const query = `
	-- beginsql
	DELETE FROM products WHERE id = ?
	-- endsql
	`
	result, err := r.db.Exec(query, id)

	if err != nil {
		panic(err)
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		panic(err)
	}

	return rowsAffected, nil
}

func (r *repository) GetBestSellers() ([]*ProductTop, error) {
	const query = `
		-- beginsql
		SELECT od.product_id,
				p.product_name,
				SUM(od.quantity*od.unit_price) vendidos 
			FROM order_details od INNER JOIN products p ON od.product_id = p.id
			GROUP BY od.product_id
			ORDER BY vendidos DESC
			LIMIT 10
		-- endsql
	`

	results, err := r.db.Query(query)

	if err != nil {
		panic(err)
	}

	var bestSellers []*ProductTop

	for results.Next() {
		product := ProductTop{}
		err = results.Scan(&product.ID, &product.ProductName, &product.Sellers)
		if err != nil {
			panic(err)
		}
		bestSellers = append(bestSellers, &product)
	}

	return bestSellers, nil
}

func (r *repository) GetTotalSellers() (float64, error) {
	const query = `
		-- beginsql
		SELECT SUM(od.quantity * od.unit_price) vendido FROM order_details OD
		-- endsql
	`
	result := r.db.QueryRow(query)
	var totalSellers float64
	err := result.Scan(&totalSellers)

	if err != nil {
		panic(err)
	}

	return totalSellers, nil
}
