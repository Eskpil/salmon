package models

type Machine struct {
	Id         string      `json:"id" bson:"_id"`
	Name       string      `json:"name" bson:"name"`
	Groups     []string    `json:"groups" bson:"groups"`
	NodeId     string      `json:"node_id", bson:"node_id"`
	Hostname   string      `json:"hostname" bson:"hostname"`
	Interfaces []Interface `json:"interfaces" bson:"interfaces"`
}
