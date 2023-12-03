package model

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

type BookingDetailRQ struct {
	BookingType        string       `json:"bookingType" bson:"booking_type"`
	BookingServiceTime *ServiceTime `json:"bookingServiceTime" bson:"booking_service_time"`
	DetailService      interface{}  `json:"detailService" bson:"detail_service"`
	CreatedAt          *time.Time   `json:"createdAt" bson:"created_at"`
	CreatedBy          *string      `json:"createdBy" bson:"created_by"`
	UpdatedAt          *time.Time   `json:"updatedAt" bson:"updated_at"`
	UpdatedBy          *string      `json:"updatedBy" bson:"updated_by"`
}

type BookingDetail struct {
	Id                 string       `json:"id,omitempty" bson:"_id"`
	BookingType        string       `json:"bookingType" bson:"booking_type"`
	BookingServiceTime *ServiceTime `json:"bookingServiceTime" bson:"booking_service_time"`
	DetailService      interface{}  `json:"detailService" bson:"detail_service"`
}

type ServiceTime struct {
	StartTime *time.Time `json:"startTime" bson:"start_time"`
	Duration  int        `json:"duration" bson:"duration"`
}

func (f ServiceTime) Value() (driver.Value, error) {
	return json.Marshal(f)
}

// Scan implements the sql.Scanner interface
func (f *ServiceTime) Scan(value interface{}) error {

	var data []byte
	switch value.(type) {
	case []uint8:
		data = []byte(value.([]uint8))
	case string:
		data = []byte(value.(string))
	}

	return json.Unmarshal(data, &f)
}

func (f BookingDetail) Value() (driver.Value, error) {
	return json.Marshal(f)
}

// Scan implements the sql.Scanner interface
func (f *BookingDetail) Scan(value interface{}) error {

	var data []byte
	switch value.(type) {
	case []uint8:
		data = []byte(value.([]uint8))
	case string:
		data = []byte(value.(string))
	}

	return json.Unmarshal(data, &f)
}
