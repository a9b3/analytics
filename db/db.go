package db

import (
	mgo "gopkg.in/mgo.v2"
)

// Init returns an instance of mgo.Database
func Init(uri, name string) (*mgo.Database, error) {
	session, err := mgo.Dial(uri)
	if err != nil {
		return nil, err
	}

	return session.DB(name), nil
}