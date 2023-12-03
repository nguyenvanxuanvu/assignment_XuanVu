package model

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

type BookingDetail struct {
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
