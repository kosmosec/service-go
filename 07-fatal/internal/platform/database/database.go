package database

import (
	"github.com/jmoiron/sqlx"
	"net/url"
	_ "github.com/lib/pq" // The database driver in use
)

type Config struct {
	User string
	Password string
	Host string
	Name string
	DisableTLS bool
}

func Open(cfg Config) (*sqlx.DB, error) {
	sslMode := "require"
	if !cfg.DisableTLS {
		sslMode = "disable"
	}
	// Query params
	q := make(url.Values)
	q.Set("sslmode", sslMode)
	q.Set("timezone", "utc")

	// construct url
	u := url.URL{
		Scheme: "postgres",
		User: url.UserPassword(cfg.User, cfg.Password),
		Host: cfg.Host,
		Path: cfg.Name,
		RawQuery: q.Encode(),
	}

	return sqlx.Open("postgres", u.String())
}
