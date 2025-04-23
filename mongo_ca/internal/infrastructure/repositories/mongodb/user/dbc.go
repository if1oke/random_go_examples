package user

import (
	"context"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"mongo_ca/internal/domain/entity/user"
)

type UserRepository struct {
	col *mongo.Collection
}

func NewUserRepo(col *mongo.Collection) *UserRepository {
	return &UserRepository{col}
}

func (db *UserRepository) GetUser(ctx context.Context, id string) (user.User, error) {
	var mongoUser MongoUser
	uid, _ := bson.ObjectIDFromHex(id)
	err := db.col.FindOne(ctx, bson.M{"_id": uid}).Decode(&mongoUser)
	if err != nil {
		return user.User{}, err
	}
	return FromMongoUser(mongoUser), nil
}

func (db *UserRepository) CreateUser(ctx context.Context, usr user.User) (user.User, error) {
	mongoUser := ToMongoUser(usr)
	_, err := db.col.InsertOne(ctx, mongoUser)
	if err != nil {
		return user.User{}, err
	}
	return FromMongoUser(mongoUser), nil
}

func (db *UserRepository) ListUser(ctx context.Context) ([]user.User, error) {
	var mongoUsers []MongoUser
	var entities []user.User
	cursor, err := db.col.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var mongoUser MongoUser
		err = cursor.Decode(&mongoUser)
		if err != nil {
			return nil, err
		}
		mongoUsers = append(mongoUsers, mongoUser)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	for _, mongoUser := range mongoUsers {
		usr := FromMongoUser(mongoUser)
		entities = append(entities, usr)
	}

	return entities, nil
}

func (db *UserRepository) FindByUsername(ctx context.Context, username string) (user.User, error) {
	var mongoUser MongoUser
	err := db.col.FindOne(ctx, bson.M{"username": username}).Decode(&mongoUser)
	if err != nil {
		return user.User{}, err
	}
	return FromMongoUser(mongoUser), nil
}
