package order

import (
	"database/sql"
	"fmt"
)

type Repository interface {
	GetOrderByID(param *getOrderByIDRequest) (*OrderItem, error)
	GetOrders(params *getOrdersRequest) ([]*OrderItem, error)
	GetTotalOrders(params *getOrdersRequest) (int64, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db}
}

func (r *repository) GetOrderByID(param *getOrderByIDRequest) (*OrderItem, error) {
	const query = `
	-- beginsql
	SELECT o.id, o.customer_id, o.order_date, o.status_id, os.status_name, CONCAT(c.first_name, ' ', c.last_name)
		AS customer_name, c.company, c.address, c.business_phone, c.city
		FROM orders o
		INNER JOIN orders_status os ON o.status_id = os.id
		INNER JOIN customers c ON o.customer_id = c.id
		WHERE o.id = ?
	-- endsql
	`
	result := r.db.QueryRow(query, param.orderID)

	orderItem := &OrderItem{}

	err := result.Scan(&orderItem.ID, &orderItem.CustomerID, &orderItem.OrderDate, &orderItem.StatusID,
		&orderItem.StatusName, &orderItem.Customer, &orderItem.Company, &orderItem.Address, &orderItem.Phone, &orderItem.City)

	if err != nil {
		panic(err)
	}

	ordersDetail, err := getOrderDetail(r, &param.orderID)

	if err != nil {
		panic(err)
	}

	orderItem.Data = ordersDetail

	return orderItem, nil
}

func getOrderDetail(repo *repository, orderId *int64) ([]*OrderDetailItem, error) {
	const query = `
	-- beginsql
	SELECT order_id, od.id, quantity, unit_price, p.product_name, product_id 
	FROM order_details od INNER JOIN products p ON od.product_id = p.id 
	WHERE od.order_id = ?
	-- endsql
	`
	results, err := repo.db.Query(query, orderId)
	if err != nil {
		panic(err)
	}
	var orderDetailItems []*OrderDetailItem
	for results.Next() {
		orderDetailItem := OrderDetailItem{}
		results.Scan(&orderDetailItem.OrderID, &orderDetailItem.ID, &orderDetailItem.Quantity, &orderDetailItem.UnitPrice,
			&orderDetailItem.ProductName, &orderDetailItem.ProductID)
		orderDetailItems = append(orderDetailItems, &orderDetailItem)
	}

	return orderDetailItems, nil
}

func (r *repository) GetOrders(params *getOrdersRequest) ([]*OrderItem, error) {
	var filter string

	if params.Status != nil {
		filter += fmt.Sprintf(" AND o.status_id = %v ", params.Status.(float64))
	}

	if params.DateFrom != nil && params.DateTo == nil {
		filter += fmt.Sprintf(" AND o.order_date >= '%v' ", params.DateFrom.(string))
	}

	if params.DateFrom == nil && params.DateTo != nil {
		filter += fmt.Sprintf(" AND o.order_date <= '%v' ", params.DateTo.(string))
	}

	if params.DateFrom != nil && params.DateTo != nil {
		filter += fmt.Sprintf(" AND o.order_date BETWEEN '%v' AND '%v' ", params.DateFrom.(string), params.DateTo.(string))
	}

	var query = `
	-- beginsql
	SELECT o.id, o.customer_id, o.order_date, o.status_id, os.status_name, CONCAT(c.first_name, ' ', c.last_name)
		AS customer_name
		FROM orders o
		INNER JOIN orders_status os ON o.status_id = os.id
		INNER JOIN customers c ON o.customer_id = c.id
		WHERE 1 = 1
	-- endsql
	`

	query = query + filter + `
	-- beginsql 
		LIMIT ? OFFSET ? 
	-- endsql
	`

	results, err := r.db.Query(query, params.Limit, params.Offset)

	if err != nil {
		panic(err)
	}
	var orders []*OrderItem
	for results.Next() {
		order := OrderItem{}
		err := results.Scan(&order.ID, &order.CustomerID, &order.OrderDate, &order.StatusID, &order.StatusName, &order.Customer)
		if err != nil {
			panic(err)
		}
		orderDetail, err := getOrderDetail(r, &order.ID)
		if err != nil {
			panic(err)
		}
		order.Data = orderDetail
		orders = append(orders, &order)
	}

	return orders, nil
}

func (r *repository) GetTotalOrders(params *getOrdersRequest) (int64, error) {
	var filter string

	if params.Status != nil {
		filter += fmt.Sprintf(" AND o.status_id = %v ", params.Status.(float64))
	}

	if params.DateFrom != nil && params.DateTo == nil {
		filter += fmt.Sprintf(" AND o.order_date >= '%v' ", params.DateFrom.(string))
	}

	if params.DateFrom == nil && params.DateTo != nil {
		filter += fmt.Sprintf(" AND o.order_date <= '%v' ", params.DateTo.(string))
	}

	if params.DateFrom != nil && params.DateTo != nil {
		filter += fmt.Sprintf(" AND o.order_date BETWEEN '%v' AND '%v' ", params.DateFrom.(string), params.DateTo.(string))
	}

	var query = `
		-- beginsql
		SELECT COUNT(id) FROM orders o WHERE 1 = 1
		-- endsql
	` + filter

	row := r.db.QueryRow(query)

	var totalRecords int64
	err := row.Scan(&totalRecords)
	if err != nil {
		panic(err)
	}

	return totalRecords, nil
}
