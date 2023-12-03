package handler

import (
	"encoding/json"
	"net/http"

	"go-service/internal/model"
	"go-service/internal/service"

	
)

func NewSendJobHandler(service service.SendJobService) *SendJobHandler {
	return &SendJobHandler{service: service}
}

type SendJobHandler struct {
	service service.SendJobService
}



func (h *SendJobHandler) SendNoti(w http.ResponseWriter, r *http.Request) {
	var rq model.NotiRequest
	er1 := json.NewDecoder(r.Body).Decode(&rq)
	defer r.Body.Close()
	if er1 != nil {
		http.Error(w, er1.Error(), http.StatusBadRequest)
		return
	}

	res, er2 := h.service.SendJob(r.Context(), rq.JobId, rq.UserId)
	if res < 0 {
		JSON(w, http.StatusNotFound, res)
		return
	}
	if er2 != nil {
		http.Error(w, er1.Error(), http.StatusInternalServerError)
		return
	}
	JSON(w, http.StatusCreated, res)
	
}

func JSON(w http.ResponseWriter, code int, res interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(res)
}
