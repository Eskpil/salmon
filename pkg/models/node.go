package models

type Node struct {
	Id       string `json:"id" bson:"_id"`
	Hostname string `json:"hostname" bson"hostname"`
}
