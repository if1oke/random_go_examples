package user

import (
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"mongo_ca/internal/infrastructure/repositories/mongodb/user"
)

type Application struct {
	client   *mongo.Client
	useCases *UseCases
}

func NewApplication() *Application {
	return &Application{}
}

func (app *Application) Client() *mongo.Client {
	if app.client == nil {
		app.client, _ = mongo.Connect(options.Client().ApplyURI("mongodb://localhost:27017"))
	}
	return app.client
}

func (app *Application) UseCases() *UseCases {
	if app.useCases == nil {
		app.useCases = NewUseCases(
			user.NewUserRepo(app.Client().Database("app").Collection("user")),
		)
	}
	return app.useCases
}
