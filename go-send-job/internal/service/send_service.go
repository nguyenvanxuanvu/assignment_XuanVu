package service

import (
	"context"
	"encoding/json"
	"fmt"
	"go-service/internal/model"
	"go-service/internal/repository"

	"firebase.google.com/go/messaging"
)

type SendJobService interface {
	SendJob(ctx context.Context, jobId string, userId []string) (int64, error)
}

func NewSendJobService(messageClient *messaging.Client, repository repository.JobRepository, detailRepository repository.BookingDetailRepository, notiRepository repository.NotificationRepository, deviceTokenRepository repository.DeviceTokenRepository) SendJobService {
	return &sendJobService{messageClient: messageClient, repository: repository, detailRepository: detailRepository, notiRepository: notiRepository, deviceTokenRepository: deviceTokenRepository}
}

type sendJobService struct {
	messageClient         *messaging.Client
	repository            repository.JobRepository
	detailRepository      repository.BookingDetailRepository
	notiRepository        repository.NotificationRepository
	deviceTokenRepository repository.DeviceTokenRepository
}

func (s *sendJobService) SendJob(ctx context.Context, jobId string, userId []string) (int64, error) {
	booking, err := s.repository.Load(ctx, jobId)
	if err != nil {
		return -1, err
	}
	if booking == nil {
		return -1, err
	}
	bookingDetail, err := s.detailRepository.Load(ctx, booking.BookingDetailId)
	if err != nil {
		return -1, err
	}
	if bookingDetail == nil {
		return -1, err
	}
	booking.BookingDetail = bookingDetail

	// get list device token
	var deviceTokens []model.TokenDevice
	deviceTokens, err = s.deviceTokenRepository.Load(ctx, userId)
	if err != nil {
		return -1, err
	}
	tokenList := getListDeviceToken(ctx, deviceTokens)

	
	// will run in transaction in the future

	// send notification
	var inInterface map[string]string
	inrec, _ := json.Marshal(booking)
	json.Unmarshal(inrec, &inInterface)

	response, err := s.messageClient.SendMulticast(context.Background(), &messaging.MulticastMessage{
		Notification: &messaging.Notification{
			Title: "New Job!!",
			Body:  "You have just received new job",
		},
		Data:   inInterface,
		Tokens: tokenList, // it's an array of device tokens
	})
	fmt.Println(response.Responses[0].Error)

	var noti model.Notification
	noti.JobId = jobId
	noti.JobDetail = bookingDetail
	noti.UserId = userId
	noti.DeviceToken = tokenList

	// insert record notification
	_, err = s.notiRepository.Create(ctx, &noti)
	if err != nil {
		return -1, err
	}

	return 1, nil
}

func getListDeviceToken(ctx context.Context, tokenDevice []model.TokenDevice) []string {
	var result []string
	for _, ele := range tokenDevice {
		result = append(result, ele.Token)
	}
	return result
}
