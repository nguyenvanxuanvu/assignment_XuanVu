package service

import (
	"context"
	"go-service/internal/model"
	"go-service/internal/repository"
)

type PriceService interface {
	Calculate(ctx context.Context, bookingDetail *model.BookingDetail) (float32, error)
}

func NewPriceService(repository repository.PriceRepository, unitPriceRepository repository.UnitPriceRepository) PriceService {
	return &priceService{repository: repository, unitPriceRepository: unitPriceRepository}
}

type priceService struct {
	repository          repository.PriceRepository
	unitPriceRepository repository.UnitPriceRepository
}

func (s *priceService) Calculate(ctx context.Context, bookingDetail *model.BookingDetail) (float32, error) {

	// get price
	bookingDate := *bookingDetail.BookingServiceTime.StartTime
	price, err := s.repository.Load(ctx, bookingDetail.BookingType, bookingDate.Format("2006-01-02"))
	if err != nil {
		return -1, err
	}
	if price == nil {
		return -1, err
	}
	// get unit price
	unitPrice, err := s.unitPriceRepository.Load(ctx, price.UnitPriceId)
	if err != nil {
		return -1, err
	}
	if unitPrice == nil {
		return -1, err
	}
	calculatedPrice := float32(bookingDetail.BookingServiceTime.Duration) / float32(unitPrice.Duration) * float32(unitPrice.Price)

	return calculatedPrice, nil

}
