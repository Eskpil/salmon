package tasks

import (
	"context"

	"github.com/eskpil/salmon/pkg/definitions"
	"github.com/eskpil/salmon/pkg/models"
	"github.com/eskpil/salmon/services/api/mycontext"
	"github.com/google/uuid"
)

func CreateTask(ctx context.Context, c *mycontext.Context, name string, nodeId string) (*definitions.Task, error) {
	task := models.Task{
		Id:     uuid.New().String(),
		NodeId: nodeId,
		Name:   name,
		Status: models.Pending,
	}

	_, err := c.Db.Collection("tasks").InsertOne(ctx, &task)

	if err != nil {
		return nil, err
	}

	return &definitions.Task{Id: task.Id, Name: name}, nil
}
