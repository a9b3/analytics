package db

// UserColName is the mongo collection name for user
const UserColName = "user"

type User struct {
	ID string `json:"_id" bson:"_id"`
}
