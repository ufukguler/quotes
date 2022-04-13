package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Version struct {
	Id      primitive.ObjectID `bson:"_id" json:"id"`
	Key     string             `bson:"key" json:"key"`
	Current int64              `bson:"current" json:"current"`
	Minimum int64              `bson:"minimum" json:"minimum"`
}

type VersionResponse struct {
	MinimumVersion int64 `json:"minimumVersion"`
	ShouldUpdate   bool  `json:"shouldUpdate"`
	ForceUpdate    bool  `json:"forceUpdate"`
}
