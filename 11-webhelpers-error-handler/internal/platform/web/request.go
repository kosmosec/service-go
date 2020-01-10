package web

import (
	"encoding/json"
	"net/http"
)

func Decode(r *http.Request, v interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		return NewRequestError(err, http.StatusBadRequest)
	}
	return nil
}
