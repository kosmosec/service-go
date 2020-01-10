package handlers

import (
	"encoding/json"
	"github.com/ardanlabs/tomhol/08-logging/internal/product"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
)

type products struct {
	db *sqlx.DB
	log *log.Logger
}

func NewProducts(db *sqlx.DB, log *log.Logger) *products {
	return &products{db: db, log: log}
}

func (p *products) List(w http.ResponseWriter, r *http.Request) {
	list, err := product.List(p.db)
	if err != nil {
		p.log.Printf("error: listing products: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(list)
	if err != nil {
		p.log.Println("error marshalling result", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(data); err != nil {
		p.log.Println("error writing result", err)
	}
}
