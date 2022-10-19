package machines

import (
	"context"

	"github.com/eskpil/salmon/pkg/models"
	"github.com/eskpil/salmon/services/api/mycontext"
)

func CreateMachine(ctx context.Context, c *mycontext.Context, machineId string, machineName string, machineGroups []string, machineHostname string, nodeId string) (*models.Machine, error) {
	// Ideally we should just pass this as the argument
	machine := &models.Machine{
		Id:       machineId,
		Name:     machineName,
		Groups:   machineGroups,
		Hostname: machineHostname,
		NodeId:   nodeId,
	}

	_, err := c.Db.Collection("machines").InsertOne(ctx, &machine)

	if err != nil {
		return nil, err
	}

	return machine, nil
}
