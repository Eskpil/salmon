package tasks

import (
	"context"

	"github.com/eskpil/salmon/pkg/models"
	"github.com/eskpil/salmon/services/api/mycontext"
	"go.mongodb.org/mongo-driver/bson"
)

func ChangeTaskStatus(ctx context.Context, c *mycontext.Context, taskId string, status models.TaskStatus) error {
	filter := bson.D{{"_id", taskId}}
	update := bson.D{{"$set", bson.D{{"status", status}}}}

	result, err := c.Db.Collection("tasks").UpdateOne(ctx, filter, update)

	if err != nil {
		return err
	}

	// TODO: We should probably do something with result
	_ = result

	return nil
}

func FinishTask(ctx context.Context, c *mycontext.Context, taskId string) error {
	return ChangeTaskStatus(ctx, c, taskId, models.Finished)
}

func FailTask(ctx context.Context, c *mycontext.Context, taskId string) error {
	return ChangeTaskStatus(ctx, c, taskId, models.Failed)
}
