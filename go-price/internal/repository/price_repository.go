package repository

import (
	"context"

	"go-service/internal/model"
)

type PriceRepository interface {
	Load(ctx context.Context, bookingType string, date string) (*model.PriceInfo, error)
}
type UnitPriceRepository interface {
	Load(ctx context.Context, id string) (*model.UnitPrice, error)
}
