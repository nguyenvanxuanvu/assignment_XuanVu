package repository

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"strings"

	"go-service/internal/model"
)

func NewPriceAdapter(db *mongo.Database) *PriceAdapter {
	return &PriceAdapter{Collection: db.Collection("prices")}
}

type PriceAdapter struct {
	Collection *mongo.Collection
}

func (r *PriceAdapter) Load(ctx context.Context, bookingType string, date string) (*model.PriceInfo, error) {

	filter := bson.M{"date": date, "booking_type": bookingType}
	res := r.Collection.FindOne(ctx, filter)
	if res.Err() != nil {
		if strings.Compare(fmt.Sprint(res.Err()), "mongo: no documents in result") == 0 {
			filter := bson.M{"date": "0000-00-00", "booking_type": bookingType}
			res = r.Collection.FindOne(ctx, filter)
			if strings.Compare(fmt.Sprint(res.Err()), "mongo: no documents in result") == 0 {
				return nil, nil
			}
		} else {
			return nil, res.Err()
		}
	}

	var price model.PriceInfo
	err := res.Decode(&price)
	fmt.Println(price)
	if err != nil {
		return nil, err
	}
	return &price, nil
}

func ToObjectId(id string) (*primitive.ObjectID, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return &objectId, nil
}
