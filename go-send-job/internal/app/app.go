package app

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/api/option"

	"go-service/internal/handler"
	"go-service/internal/repository"
	"go-service/internal/service"
)

type ApplicationContext struct {
	SendJob handler.SendJobPort
}

func NewApp(ctx context.Context, conf Config) (*ApplicationContext, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(conf.Mongo.Uri))
	db := client.Database(conf.Mongo.Database)
	if err != nil {
		return nil, err
	}

	opt := option.WithCredentialsFile("serviceAccountKey.json")
	app, err := firebase.NewApp(context.TODO(), nil, opt)
	if err != nil {
		log.Fatalf("new firebase app: %s", err)
	}

	fcmClient, err := app.Messaging(context.TODO())
	if err != nil {
		log.Fatalf("messaging: %s", err)
	}

	jobRepository := repository.NewJobAdapter(db)

	jobDetailRepository := repository.NewBookingDetailAdapter(db)

	notificationRepository := repository.NewNotificationAdapter(db)

	deviceTokenRepository := repository.NewDeviceTokenAdapter(db)

	sendJobService := service.NewSendJobService(fcmClient, jobRepository, jobDetailRepository, notificationRepository, deviceTokenRepository)
	sendJobHandler := handler.NewSendJobHandler(sendJobService)

	return &ApplicationContext{
		SendJob: sendJobHandler,
	}, nil
}
