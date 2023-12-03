package model

type TokenDevice struct {
	UserId string `json:"userId" bson:"user_id"`
	Token  string `json:"token" bson:"token"`
}
