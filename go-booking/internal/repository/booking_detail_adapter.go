package repository

import (
	"context"
	"fmt"

	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

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

func (r *BookingDetailAdapter) Create(ctx context.Context, bookingDetail *model.BookingDetailRQ) (string, error) {
	res, err := r.Collection.InsertOne(ctx, bookingDetail)
	if err != nil {
		errMsg := err.Error()
		if strings.Index(errMsg, "duplicate key error collection:") >= 0 {
			if strings.Index(errMsg, "dup key: { _id: ") >= 0 {
				return "", nil
			} else {
				return "", nil
			}
		}
		return "", err
	}
	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (r *BookingDetailAdapter) Update(ctx context.Context, bookingDetail *model.BookingDetailRQ, id string) (int64, error) {

	objectId, err := ToObjectId(id)
	if err != nil {
		return -1, err
	}

	filter := bson.M{"_id": objectId}
	update := bson.M{
		"$set": bookingDetail,
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

func (r *BookingDetailAdapter) Delete(ctx context.Context, id string) (int64, error) {
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
