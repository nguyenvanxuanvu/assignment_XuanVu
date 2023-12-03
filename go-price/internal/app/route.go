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

	price := "/prices"
	r.HandleFunc(price+"/calculate", app.Price.Calculate).Methods(POST)

	return nil
}
