package db

import (
	"gopkg.in/mgo.v2/bson"
)

// ApplicationColName is the mongo collection name for application
const ApplicationColName = "application"

type Application struct {
	ID     bson.ObjectId `json:"_id" bson:"_id", omitempty`
	Name   string        `json:"name" bson:"name"`
	UserID string        `json:"userID" bson:"userID"`
}
