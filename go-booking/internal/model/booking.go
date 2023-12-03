package model

import "time"

type BookingRQ struct {
	CustomerId       string           `json:"customerId" bson:"customer_id"`
	BookingDetail    *BookingDetailRQ `json:"bookingDetail" bson:"booking_detail"`
	Status           string           `json:"status" bson:"status"`
	BookingDetailId  string           `json:"bookingDetailId" bson:"booking_detail_id"`
	DomesticWorkerId []string         `json:"domesticWorkerId" bson:"domestic_worker_id"`
	CreatedAt        *time.Time       `json:"createdAt" bson:"created_at"`
	CreatedBy        *string          `json:"createdBy" bson:"created_by"`
	UpdatedAt        *time.Time       `json:"updatedAt" bson:"updated_at"`
	UpdatedBy        *string          `json:"updatedBy" bson:"updated_by"`
}

type Booking struct {
	Id               string         `json:"id,omitempty" bson:"_id"`
	CustomerId       string         `json:"customerId" bson:"customer_id"`
	BookingDetailId  string         `json:"bookingDetailId" bson:"booking_detail_id"`
	BookingDetail    *BookingDetail `json:"bookingDetail" bson:"booking_detail"`
	Status           string         `json:"status" bson:"status"`
	DomesticWorkerId []string       `json:"domesticWorkerId" bson:"domestic_worker_id"`
	CreatedAt        *time.Time     `json:"createdAt" bson:"created_at"`
	CreatedBy        *string        `json:"createdBy" bson:"created_by"`
	UpdatedAt        *time.Time     `json:"updatedAt" bson:"updated_at"`
	UpdatedBy        *string        `json:"updatedBy" bson:"updated_by"`
}
