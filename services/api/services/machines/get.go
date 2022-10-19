package machines

import (
	"context"

	"github.com/eskpil/salmon/pkg/models"
	"github.com/eskpil/salmon/services/api/mycontext"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CheckMachineExistance(ctx context.Context, c *mycontext.Context, machineId string) bool {
	type nothing struct{}
	var data nothing
	err := c.Db.Collection("machines").FindOne(ctx, bson.D{{"_id", machineId}}).Decode(&data)

	if err != nil && err == mongo.ErrNoDocuments {
		return false
	}

	return true
}

func GetAll(ctx context.Context, c *mycontext.Context) ([]models.Machine, error) {
	cursor, err := c.Db.Collection("machines").Find(ctx, bson.D{})

	var results []models.Machine

	if err != nil {
		return results, err
	}

	if err = cursor.All(context.TODO(), &results); err != nil {
		return results, err
	}

	return results, nil
}

func GetById(ctx context.Context, c *mycontext.Context, id string) (*models.Machine, error) {
	var machine models.Machine
	err := c.Db.Collection("machines").FindOne(ctx, bson.D{{"_id", id}}).Decode(&machine)
	if err != nil {
		return nil, err
	}

	return &machine, nil
}
