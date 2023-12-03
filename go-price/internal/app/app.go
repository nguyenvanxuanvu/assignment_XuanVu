package app

import (
	"context"
	

	"go-service/internal/handler"
	
	"go-service/internal/repository"
	"go-service/internal/service"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ApplicationContext struct {
	Price   handler.PricePort
}

func NewApp(ctx context.Context, conf Config) (*ApplicationContext, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(conf.Mongo.Uri))
	db := client.Database(conf.Mongo.Database)
	if err != nil {
		return nil, err
	}

	
	priceRepository := repository.NewPriceAdapter(db)
	
	unitPriceRepository := repository.NewUnitPriceAdapter(db)
	
	priceService := service.NewPriceService(priceRepository, unitPriceRepository)
	priceHandler := handler.NewPriceHandler(priceService)

	return &ApplicationContext{
		Price: priceHandler,
	}, nil
}
