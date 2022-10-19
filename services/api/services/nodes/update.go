package nodes

import (
	"context"
	"log"

	"github.com/eskpil/salmon/services/api/mycontext"
	"go.mongodb.org/mongo-driver/bson"
)

func SetHostname(ctx context.Context, c *mycontext.Context, nodeId string, hostname string) error {
	filter := bson.D{{"_id", nodeId}}
	update := bson.D{{"$set", bson.D{{"hostname", hostname}}}}

	result, err := c.Db.Collection("nodes").UpdateOne(ctx, filter, update)

	if err != nil {
		log.Fatal(err)
	}

	// TODO: We should probably do something with result
	_ = result

	return nil
}
