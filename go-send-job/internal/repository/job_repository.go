package repository

import (
	"context"

	"go-service/internal/model"
)

type JobRepository interface {
	Load(ctx context.Context, id string) (*model.JobInfo, error)
}
type BookingDetailRepository interface {
	Load(ctx context.Context, id string) (*model.BookingDetail, error)
}
type NotificationRepository interface {
	Create(ctx context.Context, noti *model.Notification) (*model.Notification, error)
}

type DeviceTokenRepository interface {
	Load(ctx context.Context, userId []string) ([]model.TokenDevice, error)
}
