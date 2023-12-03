package repository

import (
	"context"
	"fmt"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"go-service/internal/model"
)

func NewJobAdapter(db *mongo.Database) *JobAdapter {
	return &JobAdapter{Collection: db.Collection("bookings")}
}

type JobAdapter struct {
	Collection *mongo.Collection
}

func (r *JobAdapter) Load(ctx context.Context, id string) (*model.JobInfo, error) {
	objectId, err := ToObjectId(id)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": objectId}
	res := r.Collection.FindOne(ctx, filter)
	if res.Err() != nil {
		if strings.Compare(fmt.Sprint(res.Err()), "mongo: no documents in result") == 0 {
			return nil, nil
		} else {
			return nil, res.Err()
		}
	}
	var job model.JobInfo
	err = res.Decode(&job)
	if err != nil {
		return nil, err
	}

	return &job, nil
}
