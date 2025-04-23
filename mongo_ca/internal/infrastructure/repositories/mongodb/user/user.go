package user

import "go.mongodb.org/mongo-driver/v2/bson"

type MongoUser struct {
	ID       bson.ObjectID `bson:"_id"`
	Username string        `bson:"username"`
	Password string        `bson:"password"`
}
