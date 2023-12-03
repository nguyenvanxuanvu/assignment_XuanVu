package handler

import (
	"encoding/json"

	"net/http"

	"go-service/internal/model"

	"github.com/gorilla/mux"

	"go-service/internal/service"
)

func NewBookingHandler(service service.BookingService) *BookingHandler {
	return &BookingHandler{service: service}
}

type BookingHandler struct {
	service service.BookingService
}

func (h *BookingHandler) Load(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if len(id) == 0 {
		http.Error(w, "Id cannot be empty", http.StatusBadRequest)
		return
	}

	booking, err := h.service.Load(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	status := http.StatusOK
	if booking == nil {
		status = http.StatusNotFound
	}
	JSON(w, status, booking)
}
func (h *BookingHandler) Create(w http.ResponseWriter, r *http.Request) {
	var booking model.BookingRQ
	er1 := json.NewDecoder(r.Body).Decode(&booking)
	defer r.Body.Close()
	if er1 != nil {
		http.Error(w, er1.Error(), http.StatusBadRequest)
		return
	}

	res, er2 := h.service.Create(r.Context(), &booking)
	if er2 != nil {
		http.Error(w, er1.Error(), http.StatusInternalServerError)
		return
	}
	JSON(w, http.StatusCreated, res)
}
func (h *BookingHandler) Update(w http.ResponseWriter, r *http.Request) {
	var booking model.BookingRQ
	er1 := json.NewDecoder(r.Body).Decode(&booking)
	defer r.Body.Close()
	if er1 != nil {
		http.Error(w, er1.Error(), http.StatusBadRequest)
		return
	}
	id := mux.Vars(r)["id"]
	if len(id) == 0 {
		http.Error(w, "Id cannot be empty", http.StatusBadRequest)
		return
	}

	res, er2 := h.service.Update(r.Context(), &booking, id)
	if er2 != nil {
		http.Error(w, er2.Error(), http.StatusInternalServerError)
		return
	}
	if res <= 0 {
		JSON(w, http.StatusNotFound, res)
	} else {
		JSON(w, http.StatusOK, res)
	}
}
func (h *BookingHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if len(id) == 0 {
		http.Error(w, "Id cannot be empty", http.StatusBadRequest)
		return
	}
	res, err := h.service.Delete(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if res <= 0 {
		JSON(w, http.StatusNotFound, res)
	} else {
		JSON(w, http.StatusOK, res)
	}
}

func JSON(w http.ResponseWriter, code int, res interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(res)
}
