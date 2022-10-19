package tasks

import (
	"context"

	"github.com/eskpil/salmon/pkg/models"
	"github.com/eskpil/salmon/services/api/mycontext"
	"go.mongodb.org/mongo-driver/bson"
)

func GetTaskById(ctx context.Context, c *mycontext.Context, taskId string) (*models.Task, error) {
	var task models.Task
	err := c.Db.Collection("tasks").FindOne(ctx, bson.D{{"_id", taskId}}).Decode(&task)
	if err != nil {
		return nil, err
	}

	return &task, nil
}
