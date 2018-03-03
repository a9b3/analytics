package mongo

import (
	mgo "gopkg.in/mgo.v2"
)

var db *mgo.Database

// Init sets package wide db reference
func Init(uri, name string) *mgo.Database {
	session, err := mgo.Dial(uri)
	if err != nil {
		panic(err)
	}

	db = session.DB(name)
	return db
}
