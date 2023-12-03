package handler

import (
	"encoding/json"
	"go-service/internal/model"
	"go-service/internal/service"
	"net/http"
)

func NewPriceHandler(service service.PriceService) *PriceHandler {
	return &PriceHandler{service: service}
}

type PriceHandler struct {
	service service.PriceService
}

func (h *PriceHandler) Calculate(w http.ResponseWriter, r *http.Request) {
	var bookingDetail model.BookingDetail
	er1 := json.NewDecoder(r.Body).Decode(&bookingDetail)
	defer r.Body.Close()
	if er1 != nil {
		http.Error(w, er1.Error(), http.StatusBadRequest)
		return
	}

	res, er2 := h.service.Calculate(r.Context(), &bookingDetail)
	if er2 != nil {
		http.Error(w, er1.Error(), http.StatusInternalServerError)
		return
	}
	JSON(w, http.StatusOK, res)
}

func JSON(w http.ResponseWriter, code int, res interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(res)
}
