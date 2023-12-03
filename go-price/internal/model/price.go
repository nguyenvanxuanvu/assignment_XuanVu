package model

type PriceInfo struct {
	Id          string `json:"id" bson:"_id"`
	BookingType string `json:"bookingType" bson:"booking_type"`
	Date        string `json:"date" bson:"date"`
	UnitPriceId string `json:"unitPriceId" bson:"unit_price_id"`
}

type UnitPrice struct {
	Id       string `json:"id" bson:"_id"`
	Duration int    `json:"duration" bson:"duration"`
	Price    int    `json:"price" bson:"price"`
}
