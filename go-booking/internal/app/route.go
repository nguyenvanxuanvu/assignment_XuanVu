package app

import (
	"context"
	. "github.com/core-go/core"
	"github.com/gorilla/mux"
)

func Route(r *mux.Router, ctx context.Context, conf Config) error {
	app, err := NewApp(ctx, conf)
	if err != nil {
		return err
	}

	booking := "/bookings"
	r.HandleFunc(booking+"/{id}", app.Booking.Load).Methods(GET)
	r.HandleFunc(booking, app.Booking.Create).Methods(POST)
	r.HandleFunc(booking+"/{id}", app.Booking.Update).Methods(PUT)
	r.HandleFunc(booking+"/{id}", app.Booking.Delete).Methods(DELETE)

	return nil
}
