package database

import (
	"github.com/jmoiron/sqlx"
	"net/url"
	_ "github.com/lib/pq" // The database driver in use
)

func Open() (*sqlx.DB, error) {
	// Query params
	q := make(url.Values)
	q.Set("sslmode", "disable")
	q.Set("timezone", "utc")

	// construct url
	u := url.URL{
		Scheme: "postgres",
		User: url.UserPassword("postgres", "postgres"),
		Host: "localhost",
		Path: "postgres",
		RawQuery: q.Encode(),
	}

	return sqlx.Open("postgres", u.String())
}
