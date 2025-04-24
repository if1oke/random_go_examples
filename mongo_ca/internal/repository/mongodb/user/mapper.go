package user

import (
	"go.mongodb.org/mongo-driver/v2/bson"
	"mongo_ca/internal/domain/entity/user"
)

func FromMongoUser(u MongoUser) user.User {
	return user.User{
		ID:       u.ID.String(),
		Username: u.Username,
		Password: u.Password,
	}
}

func ToMongoUser(u user.User) MongoUser {
	id, _ := bson.ObjectIDFromHex(u.ID)
	return MongoUser{
		ID:       id,
		Username: u.Username,
		Password: u.Password,
	}
}
