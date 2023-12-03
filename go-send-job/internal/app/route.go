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

	sendJob := "/jobs"
	r.HandleFunc(sendJob+"/send", app.SendJob.SendNoti).Methods(POST)

	return nil
}
