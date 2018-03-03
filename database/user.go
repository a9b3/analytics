package database

import (
	mgo "gopkg.in/mgo.v2"
)

type User struct {
	ID string `json:"_id" bson:"_id"`
}

// NewUserStore returns instance of UserStore
func NewUserStore(db *mgo.Database) *UserStore {
	return &UserStore{
		collection: db.C("user"),
	}
}

// UserStore is the public api for user
type UserStore struct {
	collection *mgo.Collection
}

// Get does a Find on the user collection
func (u *UserStore) Get(q interface{}) ([]User, error) {
	results := []User{}
	if err := u.collection.Find(q).All(&results); err != nil {
		return nil, err
	}
	return results, nil
}

// Create does a Insert on the user collection
func (u *UserStore) Create(user *User) error {
	if err := u.collection.Insert(user); err != nil {
		return err
	}
	return nil
}
