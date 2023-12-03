package model

import "time"

type Notification struct {
	JobId       string         `json:"jobId" bson:"job_id"`
	JobDetail   *BookingDetail `json:"jobDetail" bson:"job_detail"`
	UserId      []string       `json:"userId" bson:"user_id"`
	DeviceToken []string       `json:"deviceToken" bson:"device_token"`
	CreatedAt   *time.Time     `json:"createdAt" bson:"created_at"`
	CreatedBy   *string        `json:"createdBy" bson:"created_by"`
	UpdatedAt   *time.Time     `json:"updatedAt" bson:"updated_at"`
	UpdatedBy   *string        `json:"updatedBy" bson:"updated_by"`
}
