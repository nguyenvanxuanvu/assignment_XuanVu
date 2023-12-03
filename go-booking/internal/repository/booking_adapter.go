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

func NewBookingAdapter(db *mongo.Database) *BookingAdapter {
	return &BookingAdapter{Collection: db.Collection("bookings")}
}

type BookingAdapter struct {
	Collection *mongo.Collection
}

func (r *BookingAdapter) Load(ctx context.Context, id string) (*model.Booking, error) {
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
	var booking model.Booking
	err = res.Decode(&booking)
	if err != nil {
		return nil, err
	}
	return &booking, nil
}

func (r *BookingAdapter) Create(ctx context.Context, booking *model.BookingRQ) (*model.BookingRQ, error) {
	_, err := r.Collection.InsertOne(ctx, booking)
	if err != nil {
		errMsg := err.Error()
		if strings.Index(errMsg, "duplicate key error collection:") >= 0 {
			if strings.Index(errMsg, "dup key: { _id: ") >= 0 {
				return nil, nil
			} else {
				return nil, nil
			}
		}
		return nil, err
	}
	return booking, nil
}

func (r *BookingAdapter) Update(ctx context.Context, booking *model.BookingRQ, id string) (int64, error) {
	objectId, err := ToObjectId(id)
	if err != nil {
		return -1, err
	}
	filter := bson.M{"_id": objectId}
	update := bson.M{
		"$set": booking,
	}
	res, err := r.Collection.UpdateOne(ctx, filter, update)
	if res.ModifiedCount > 0 {
		return res.ModifiedCount, err
	} else if res.UpsertedCount > 0 {
		return res.UpsertedCount, err
	} else {
		return res.MatchedCount, err
	}
}

func (r *BookingAdapter) Delete(ctx context.Context, id string) (int64, error) {
	objectId, err := ToObjectId(id)
	if err != nil {
		return -1, err
	}
	filter := bson.M{"_id": objectId}
	res, err := r.Collection.DeleteOne(ctx, filter)
	if res == nil || err != nil {
		return 0, err
	}
	return res.DeletedCount, err
}

func ToObjectId(id string) (*primitive.ObjectID, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return &objectId, nil
}
