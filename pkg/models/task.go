package models

type Task struct {
	Id string `json:"id" bson:"_id"`

	// The id of the node supposed to perform this task.
	NodeId string `json:"node_id" bson:"node_id"`

	Name   string     `json:"name" bson:"name"`
	Status TaskStatus `json:"status" bson:"status"`
}

type TaskStatus int64

const (
	Pending TaskStatus = iota
	Finished
	Failed
)
