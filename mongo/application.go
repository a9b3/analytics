package mongo

import "gopkg.in/mgo.v2/bson"

var applicationCollection = "application"

type Application struct {
	Id bson.ObjectId `json:"id" bson:"_id"`
}

func CreateApplication() Application {
	a := Application{
		Id: bson.NewObjectId(),
	}
	return a
}
