package service

import (
	"context"

	"go-service/internal/model"
	"go-service/internal/repository"
)

type BookingService interface {
	Load(ctx context.Context, id string) (*model.Booking, error)
	Create(ctx context.Context, booking *model.BookingRQ) (*model.BookingRQ, error)
	Update(ctx context.Context, booking *model.BookingRQ, id string) (int64, error)
	Delete(ctx context.Context, id string) (int64, error)
}

func NewBookingService(repository repository.BookingRepository, detailRepository repository.BookingDetailRepository) BookingService {
	return &bookingService{repository: repository, detailRepository: detailRepository}
}

type bookingService struct {
	repository       repository.BookingRepository
	detailRepository repository.BookingDetailRepository
}

func (s *bookingService) Load(ctx context.Context, id string) (*model.Booking, error) {
	booking, err := s.repository.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	bookingDetail, err := s.detailRepository.Load(ctx, booking.BookingDetailId)
	if err != nil {
		return nil, err
	}
	res, err := s.repository.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	res.BookingDetail = bookingDetail
	return res, nil
}
func (s *bookingService) Create(ctx context.Context, booking *model.BookingRQ) (*model.BookingRQ, error) {

	bookingDetailId, err := s.detailRepository.Create(ctx, booking.BookingDetail)
	if err != nil {
		return nil, err
	}
	booking.BookingDetailId = bookingDetailId
	bookingRes, err := s.repository.Create(ctx, booking)
	if err != nil {
		return nil, err
	}
	return bookingRes, nil

}
func (s *bookingService) Update(ctx context.Context, booking *model.BookingRQ, id string) (int64, error) {
	existedBooking, err := s.repository.Load(ctx, id)
	if err != nil {
		return -1, err
	}

	_, err = s.detailRepository.Update(ctx, booking.BookingDetail, existedBooking.BookingDetailId)
	if err != nil {
		return -1, err
	}
	return s.repository.Update(ctx, booking, id)
}

func (s *bookingService) Delete(ctx context.Context, id string) (int64, error) {
	booking, err := s.repository.Load(ctx, id)
	if err != nil {
		return -1, err
	}
	if booking == nil {
		return -1, err
	}
	_, err = s.detailRepository.Delete(ctx, booking.BookingDetailId)
	if err != nil {
		return -1, err
	}
	return s.repository.Delete(ctx, id)
}
