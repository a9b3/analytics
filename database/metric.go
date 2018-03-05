package database

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Metric struct {
	ID    bson.ObjectId          `json:"_id" bson:"_id", omitempty`
	Name  string                 `json:"name" bson:"name"`
	AppID string                 `json:"app_id" bson:"app_id"`
	Data  map[string]interface{} `json:"data" bson:"data"`
}

// NewMetricStore returns instance of MetricStore
func NewMetricStore(db *mgo.Database) *MetricStore {
	return &MetricStore{
		collection: db.C("metric"),
	}
}

// MetricStore is the public api for metric
type MetricStore struct {
	collection *mgo.Collection
}

// Get does a Find on the metric collection
func (m *MetricStore) Get(q interface{}) ([]Metric, error) {
	results := []Metric{}
	if err := m.collection.Find(q).All(&results); err != nil {
		return nil, err
	}
	return results, nil
}

// Create does a Insert on the metric collection
func (m *MetricStore) Create(metric *Metric) error {
	metric.ID = bson.NewObjectId()
	if err := m.collection.Insert(metric); err != nil {
		return err
	}
	return nil
}
