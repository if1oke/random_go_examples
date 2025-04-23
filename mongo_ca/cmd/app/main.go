package app

import (
	"mongo_ca/internal/application/user"
	user2 "mongo_ca/internal/domain/entity/user"
)

func main() {

	userApplication := user.NewApplication()

	userApplication.UseCases().CreateUser(user2.User{
		Username: "test",
		Password: "test",
	})
}
