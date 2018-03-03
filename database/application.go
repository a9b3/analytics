package database

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Application struct {
	ID bson.ObjectId `json:"_id" bson:"_id", omitempty`
}

// NewApplicationStore returns instance of ApplicationStore
func NewApplicationStore(db *mgo.Database) *ApplicationStore {
	return &ApplicationStore{
		collection: db.C("application"),
	}
}

// ApplicationStore is the public api for application
type ApplicationStore struct {
	collection *mgo.Collection
}

// Get does a Find on the application collection
func (a *ApplicationStore) Get(q interface{}) ([]Application, error) {
	results := []Application{}
	if err := a.collection.Find(q).All(&results); err != nil {
		return nil, err
	}
	return results, nil
}

// Create does a Insert on the application collection
func (a *ApplicationStore) Create(application *Application) error {
	if err := a.collection.Insert(application); err != nil {
		return err
	}
	return nil
}
