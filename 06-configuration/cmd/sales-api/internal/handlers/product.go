package handlers

import (
	"encoding/json"
	"github.com/ardanlabs/tomhol/06-configuration/internal/product"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
)

type products struct {
	db *sqlx.DB
}

func NewProducts(db *sqlx.DB) *products {
	return &products{db: db}
}

func (p *products) List(w http.ResponseWriter, r *http.Request) {
	list, err := product.List(p.db)
	if err != nil {
		log.Printf("error: listing products: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(list)
	if err != nil {
		log.Println("error marshalling result", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(data); err != nil {
		log.Println("error writing result", err)
	}
}
