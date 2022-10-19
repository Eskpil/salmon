package models

type IpAddr struct {
	Type   int32  `json:"type" bson:"type"`
	Addr   string `json:"addr" bson:"addr"`
	Prefix uint32 `json:"prefix" bson:"prefix"`
}

type Interface struct {
	Id      string   `json:"id" bson:"_id"`
	Name    string   `json:"name" bson:"name"`
	Mac     string   `json:"mac" bson:"mac"`
	IpAddrs []IpAddr `json:"addrs" bson:"ip_addrs"`
}
