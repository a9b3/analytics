package database

import (
	mgo "gopkg.in/mgo.v2"
)

// Init sets package wide db reference
func Init(uri, name string) (*mgo.Database, error) {
	session, err := mgo.Dial(uri)
	if err != nil {
		return nil, err
	}

	return session.DB(name), nil
}
