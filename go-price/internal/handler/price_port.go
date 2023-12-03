package handler

import "net/http"

type PricePort interface {
	Calculate(w http.ResponseWriter, r *http.Request)
}
