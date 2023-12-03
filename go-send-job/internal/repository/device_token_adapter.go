package repository

import (
	"context"
	"go-service/internal/model"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewDeviceTokenAdapter(db *mongo.Database) *BookingAdapter {
	return &BookingAdapter{Collection: db.Collection("device_tokens")}
}

type BookingAdapter struct {
	Collection *mongo.Collection
}

func (r *BookingAdapter) Load(ctx context.Context, userId []string) ([]model.TokenDevice, error) {
	var results []model.TokenDevice
	filter := bson.M{"user_id": bson.M{"$in": userId}}
	cur, err := r.Collection.Find(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	err = cur.All(context.TODO(), &results)
	cur.Close(context.TODO())
	return results, nil
}

func ToObjectId(id string) (*primitive.ObjectID, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return &objectId, nil
}
