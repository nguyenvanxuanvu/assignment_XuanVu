package app

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go-service/internal/handler"
	"go-service/internal/repository"
	"go-service/internal/service"
)

type ApplicationContext struct {
	Booking handler.BookingPort
}

func NewApp(ctx context.Context, conf Config) (*ApplicationContext, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(conf.Mongo.Uri))
	db := client.Database(conf.Mongo.Database)
	if err != nil {
		return nil, err
	}

	bookingRepository := repository.NewBookingAdapter(db)

	bookingDetailRepository := repository.NewBookingDetailAdapter(db)

	bookingService := service.NewBookingService(bookingRepository, bookingDetailRepository)
	bookingHandler := handler.NewBookingHandler(bookingService)

	return &ApplicationContext{
		Booking: bookingHandler,
	}, nil
}
