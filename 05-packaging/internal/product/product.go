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
