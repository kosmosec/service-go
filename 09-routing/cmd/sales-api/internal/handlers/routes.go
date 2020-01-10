package handlers

import (
	"github.com/ardanlabs/tomhol/09-routing/internal/platform/web"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
)

func API(db *sqlx.DB, log *log.Logger) http.Handler {
	app := web.NewApp(log)

	p := NewProducts(db, log)

	app.Handle(http.MethodGet, "/v1/products", p.List)
	app.Handle(http.MethodGet, "/v1/products/{id}", p.Retrieve)

	return app
}

