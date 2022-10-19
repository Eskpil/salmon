package nodes

import (
	"context"

	"github.com/eskpil/salmon/pkg/models"
	"github.com/eskpil/salmon/services/api/mycontext"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAll(ctx context.Context, c *mycontext.Context) ([]models.Node, error) {
	cursor, err := c.Db.Collection("nodes").Find(ctx, bson.D{})

	var results []models.Node

	if err != nil {
		return results, err
	}

	if err = cursor.All(context.TODO(), &results); err != nil {
		return results, err
	}

	return results, nil
}

func GetById(ctx context.Context, c *mycontext.Context, id string) (*models.Node, error) {
	var node models.Node
	err := c.Db.Collection("nodes").FindOne(ctx, bson.D{{"_id", id}}).Decode(&node)
	if err != nil {
		return nil, err
	}

	return &node, nil
}
