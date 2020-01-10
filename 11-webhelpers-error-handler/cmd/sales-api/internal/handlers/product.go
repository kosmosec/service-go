package handlers

import (
	"github.com/ardanlabs/tomhol/11-webhelpers-error-handler/internal/platform/web"
	"github.com/ardanlabs/tomhol/11-webhelpers-error-handler/internal/product"
	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"log"
	"net/http"
	"time"
)

type products struct {
	db *sqlx.DB
	log *log.Logger
}

func NewProducts(db *sqlx.DB, log *log.Logger) *products {
	return &products{db: db, log: log}
}

func (p *products) List(w http.ResponseWriter, r *http.Request) error {
	list, err := product.List(p.db)
	if err != nil {
		return errors.Wrap(err, "getting product list")
	}

	return web.Respond(w, list, http.StatusOK)
}

func (p *products) Retrieve(w http.ResponseWriter, r *http.Request) error {
	id := chi.URLParam(r, "id")

	prod, err := product.Retrieve(p.db, id)
	if err != nil {
		switch err {
		case product.ErrNotFound:
			return web.NewRequestError(err, http.StatusNotFound)
		case product.ErrInvalidID:
			return web.NewRequestError(err, http.StatusBadRequest)
		default:
			return errors.Wrapf(err, "getting product %q", id)
		}
	}
	return web.Respond(w, prod, http.StatusOK)
}

func (p *products) Create(w http.ResponseWriter, r *http.Request) error {
	var np product.NewProduct
	if err := web.Decode(r, &np); err != nil {
		return errors.Wrap(err, "decoding new product")
	}

	prod, err := product.Create(p.db, np, time.Now())
	if err != nil {
		return errors.Wrap(err, "creating new product")
	}

	return web.Respond(w, &prod, http.StatusCreated)
}
