package db

import (
	"gopkg.in/mgo.v2/bson"
)

// MetricColName is the mongo collection name for metric
const MetricColName = "metric"

type Metric struct {
	ID    bson.ObjectId          `json:"_id" bson:"_id", omitempty`
	Name  string                 `json:"name" bson:"name"`
	AppID string                 `json:"app_id" bson:"app_id"`
	Data  map[string]interface{} `json:"data" bson:"data"`
}
