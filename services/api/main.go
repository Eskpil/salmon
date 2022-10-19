package main

import (
	"context"
	"encoding/json"
	"net"
	"sync"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/eskpil/salmon/pkg/definitions"
	"github.com/eskpil/salmon/pkg/models"
	"github.com/eskpil/salmon/services/api/mycontext"
	"google.golang.org/grpc"

	machineService "github.com/eskpil/salmon/services/api/services/machines"
	nodeService "github.com/eskpil/salmon/services/api/services/nodes"
	taskService "github.com/eskpil/salmon/services/api/services/tasks"

	machineController "github.com/eskpil/salmon/services/api/controllers/machines"
	nodeController "github.com/eskpil/salmon/services/api/controllers/nodes"
)

type M2M struct {
	definitions.UnimplementedM2MServer
	Ctx *mycontext.Context
}

func (s *M2M) Heartbeat(ctx context.Context, request *definitions.HeartbeatRequest) (*definitions.HeartbeatResponse, error) {
	log.Infof("Got heartbeat from machine: %s\n", request.MachineId)

	var node models.Node

	tasks := []*definitions.Task{}

	err := s.Ctx.Db.Collection("nodes").FindOne(ctx, bson.D{{"_id", request.MachineId}}).Decode(&node)

	if err != nil && err == mongo.ErrNoDocuments {
		log.Infof("Machine does not exist. Thus creating it")

		node.Id = request.MachineId
		node.Hostname = "pending"
		_, err := s.Ctx.Db.Collection("nodes").InsertOne(ctx, &node)
		if err != nil {
			return nil, err
		}

		log.Infof("Creating tasks")

		{
			task, err := taskService.CreateTask(ctx, s.Ctx, "collectHostname", node.Id)
			if err != nil {
				return nil, err
			}
			tasks = append(tasks, task)
		}

		{
			task, err := taskService.CreateTask(ctx, s.Ctx, "collectMachines", node.Id)
			if err != nil {
				return nil, err
			}
			tasks = append(tasks, task)
		}

		{
			task, err := taskService.CreateTask(ctx, s.Ctx, "collectInterfaces", node.Id)
			if err != nil {
				return nil, err
			}
			tasks = append(tasks, task)
		}

		{
			task, err := taskService.CreateTask(ctx, s.Ctx, "collectMachineInterfaces", node.Id)
			if err != nil {
				return nil, err
			}
			tasks = append(tasks, task)
		}
	} else if err != nil {
		log.Errorf("Error: %v\n", err)
		return nil, err
	} else {
		{
			task, err := taskService.CreateTask(ctx, s.Ctx, "collectMachines", node.Id)
			if err != nil {
				return nil, err
			}
			tasks = append(tasks, task)
		}
		{
			task, err := taskService.CreateTask(ctx, s.Ctx, "collectMachineInterfaces", node.Id)
			if err != nil {
				return nil, err
			}
			tasks = append(tasks, task)
		}

	}

	return &definitions.HeartbeatResponse{Tasks: tasks}, nil
}

func (s *M2M) FinishTask(ctx context.Context, request *definitions.FinishTaskRequest) (*definitions.FinishTaskResponse, error) {
	log.Infof("Task: %s with data: \n\n%s\n\n", request.Id, string(request.Data))
	task, err := taskService.GetTaskById(ctx, s.Ctx, request.Id)

	if err != nil {
		log.Errorf("Failed to fetch task: %s: %v\n", request.Id, err)
		// We will at least try to mark the status of the tasks as failed.
		taskService.FailTask(ctx, s.Ctx, request.Id)
		return nil, err
	}

	switch {
	case task.Name == "collectHostname":
		{
			type incoming struct {
				Hostname string `json:"hostname"`
			}

			var data incoming
			json.Unmarshal(request.GetData(), &data)
			nodeService.SetHostname(ctx, s.Ctx, task.NodeId, data.Hostname)
		}
	case task.Name == "collectMachines":
		{
			type incoming struct {
				Id       string   `json:"id"`
				Name     string   `json:"name"`
				Groups   []string `json:"groups"`
				Hostname string   `json:"hostname"`
			}

			var data []incoming
			json.Unmarshal(request.GetData(), &data)

			for _, m := range data {
				if !machineService.CheckMachineExistance(ctx, s.Ctx, m.Id) {
					machineService.CreateMachine(ctx, s.Ctx, m.Id, m.Name, m.Groups, m.Hostname, task.NodeId)
				} else {
					machineService.UpdateMachine(ctx, s.Ctx, m.Id, m.Name, m.Groups, m.Hostname)
				}
			}
		}
	case task.Name == "collectInterfaces":
		{

		}
	case task.Name == "collectMachineInterfaces":
		{
			type incoming struct {
				MachineId  string             `json:"machine_id"`
				Interfaces []models.Interface `json:"interfaces"`
			}

			var data []incoming
			json.Unmarshal(request.GetData(), &data)

			for _, i := range data {
				machineService.SetInterfaces(ctx, s.Ctx, i.MachineId, i.Interfaces)
			}
		}
	}

	taskService.FinishTask(ctx, s.Ctx, task.Id)

	return &definitions.FinishTaskResponse{}, nil
}

func main() {
	ctx, err := mycontext.NewContext()
	if err != nil {
		log.Fatalf("Failed to create a new context: %v\n", err)
	}

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()

		listener, err := net.Listen("tcp", "0.0.0.0:8080")
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		s := grpc.NewServer()

		m2m := M2M{
			Ctx: ctx,
		}

		definitions.RegisterM2MServer(s, &m2m)
		log.Printf("server listening at %v", listener.Addr())

		if err := s.Serve(listener); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}(wg)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()

		r := gin.Default()

		r.Use(func(c *gin.Context) {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
			c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, ResponseType, accept, origin, Cache-Control, X-Requested-With")
			c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

			if c.Request.Method == "OPTIONS" {
				c.AbortWithStatus(204)
				return
			}
			c.Next()
		})

		r.GET("/api/machines/", machineController.GetAll(ctx))
		r.GET("/api/machines/:id/", machineController.GetById(ctx))
		r.GET("/api/nodes/:id/", nodeController.GetById(ctx))
		r.GET("/api/nodes/", nodeController.GetAll(ctx))

		r.Run("0.0.0.0:8090")
	}(wg)

	wg.Wait()
}
