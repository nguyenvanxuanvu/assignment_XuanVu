package repository

import (
	"context"
	"strings"

	"go.mongodb.org/mongo-driver/mongo"

	"go-service/internal/model"
)

func NewNotificationAdapter(db *mongo.Database) *NotificationAdapter {
	return &NotificationAdapter{Collection: db.Collection("notifications")}
}

type NotificationAdapter struct {
	Collection *mongo.Collection
}

func (r *NotificationAdapter) Create(ctx context.Context, noti *model.Notification) (*model.Notification, error) {
	_, err := r.Collection.InsertOne(ctx, noti)
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
	return noti, nil
}
