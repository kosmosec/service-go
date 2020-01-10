package product

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

// List gets all Products from the database
func List(db *sqlx.DB) ([]Product, error) {
	products := make([]Product, 0, 10)

	const query = `SELECT * FROM products`

	if err := db.Select(&products, query); err != nil {
		return nil, errors.Wrap(err, "selecting products")
	}

	return products, nil
}

func Retrieve(db *sqlx.DB, id string) (*Product, error) {
	var p Product
	const q = `select * from products where product_id = $1`
	if err := db.Get(&p, q, id); err != nil {
		return nil, errors.Wrap(err, "selecting single products")
	}
	return &p, nil
}
