package model

import "time"

type JobInfo struct {
	Id              string         `json:"id,omitempty" bson:"_id"`
	Date            *time.Time     `json:"date" bson:"date"`
	Duration        int            `json:"duration" bson:"duration"`
	BookingDetailId string         `json:"bookingDetailId" bson:"booking_detail_id"`
	BookingDetail   *BookingDetail `json:"bookingDetail" bson:"booking_detail"`
	Status          string         `json:"status" bson:"status"`
}
