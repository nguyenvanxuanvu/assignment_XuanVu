package repository

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"strings"

	"go-service/internal/model"
)

func NewBookingDetailAdapter(db *mongo.Database) *BookingDetailAdapter {
	return &BookingDetailAdapter{Collection: db.Collection("bookings_detail")}
}

type BookingDetailAdapter struct {
	Collection *mongo.Collection
}

func (r *BookingDetailAdapter) Load(ctx context.Context, id string) (*model.BookingDetail, error) {
	objectId, err := ToObjectId(id)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": objectId}
	res := r.Collection.FindOne(ctx, filter)
	if res.Err() != nil {
		if strings.Compare(fmt.Sprint(res.Err()), "mongo: no documents in result") == 0 {
			return nil, nil
		} else {
			return nil, res.Err()
		}
	}
	var bookingDetail model.BookingDetail
	err = res.Decode(&bookingDetail)
	if err != nil {
		return nil, err
	}
	return &bookingDetail, nil
}
