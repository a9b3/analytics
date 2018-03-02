package mongo

var userCollectionName = "user"

type User struct {
	ID string `json:"id" bson:"_id"`
}

// CreateUser will insert a user into db
func CreateUser(id string) User {
	u := User{
		ID: id,
	}
	err := db.C(userCollectionName).Insert(u)
	if err != nil {
		panic(err)
	}

	return u
}
