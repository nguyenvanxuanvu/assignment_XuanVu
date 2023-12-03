package model

type NotiRequest struct {
	JobId       string         `json:"jobId" bson:"job_id"`
	UserId      []string       `json:"userId" bson:"user_id"`
}