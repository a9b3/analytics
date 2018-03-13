package db

import (
	"github.com/Sirupsen/logrus"
	mgo "gopkg.in/mgo.v2"
)

// Init returns an instance of mgo.Database
func Init(uri, name string) (*mgo.Database, error) {
	session, err := mgo.Dial(uri)
	if err != nil {
		return nil, err
	}

	db := session.DB(name)
	setCollectionInfo(db)
	return db, nil
}

func setCollectionInfo(mdb *mgo.Database) {
	index := mgo.Index{
		Key:    []string{"name", "_id"},
		Unique: true,
	}
	err := mdb.C(ApplicationColName).EnsureIndex(index)
	if err != nil {
		logrus.WithError(err).Fatalf("ensure index on application collection %s", err.Error())
	}
}
