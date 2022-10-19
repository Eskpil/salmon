package machines

import (
	"context"

	log "github.com/sirupsen/logrus"

	"github.com/eskpil/salmon/pkg/models"
	"github.com/eskpil/salmon/services/api/mycontext"
	"go.mongodb.org/mongo-driver/bson"
)

func UpdateMachine(ctx context.Context, c *mycontext.Context, machineId string, machineName string, machineGroups []string, machineHostname string) error {
	filter := bson.D{{"_id", machineId}}
	update := bson.D{{"$set", bson.D{{"hostname", machineHostname}, {"name", machineName}, {"groups", machineGroups}}}}

	result, err := c.Db.Collection("machines").UpdateOne(ctx, filter, update)

	if err != nil {
		log.Error(err)
		return err
	}

	// TODO: We should probably do something with result
	_ = result

	return nil
}

func SetInterfaces(ctx context.Context, c *mycontext.Context, machineId string, i []models.Interface) error {
	filter := bson.D{{"_id", machineId}}
	update := bson.D{{"$set", bson.D{{"interfaces", i}}}}

	result, err := c.Db.Collection("machines").UpdateOne(ctx, filter, update)

	if err != nil {
		log.Error(err)
		return err
	}

	// TODO: We should probably do something with result
	_ = result

	return nil
}
