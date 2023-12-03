package repository

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"strings"

	"go-service/internal/model"
)

func NewUnitPriceAdapter(db *mongo.Database) *UnitPriceAdapter {
	return &UnitPriceAdapter{Collection: db.Collection("unit_prices")}
}

type UnitPriceAdapter struct {
	Collection *mongo.Collection
}

func (r *UnitPriceAdapter) Load(ctx context.Context, id string) (*model.UnitPrice, error) {
	filter := bson.M{"_id": id}
	res := r.Collection.FindOne(ctx, filter)
	if res.Err() != nil {
		if strings.Compare(fmt.Sprint(res.Err()), "mongo: no documents in result") == 0 {
			return nil, nil
		} else {
			return nil, res.Err()
		}
	}
	var unitPrice model.UnitPrice
	err := res.Decode(&unitPrice)
	if err != nil {
		return nil, err
	}
	return &unitPrice, nil
}
