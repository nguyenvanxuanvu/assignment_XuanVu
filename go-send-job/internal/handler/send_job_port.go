package handler

import "net/http"

type SendJobPort interface {
	SendNoti(w http.ResponseWriter, r *http.Request)
}
