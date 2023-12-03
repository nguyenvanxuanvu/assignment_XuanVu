package repository

import (
	"context"

	"go-service/internal/model"
)

type BookingRepository interface {
	Load(ctx context.Context, id string) (*model.Booking, error)
	Create(ctx context.Context, booking *model.BookingRQ) (*model.BookingRQ, error)
	Update(ctx context.Context, booking *model.BookingRQ, id string) (int64, error)
	Delete(ctx context.Context, id string) (int64, error)
}
type BookingDetailRepository interface {
	Load(ctx context.Context, id string) (*model.BookingDetail, error)
	Create(ctx context.Context, bookingDetail *model.BookingDetailRQ) (string, error)
	Update(ctx context.Context, bookingDetail *model.BookingDetailRQ, id string) (int64, error)
	Delete(ctx context.Context, id string) (int64, error)
}
