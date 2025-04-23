package user

import (
	"context"
	"mongo_ca/internal/domain/entity/user"
	repo "mongo_ca/internal/infrastructure/repositories/mongodb/user"
)

type UseCases struct {
	repo *repo.UserRepository
}

func NewUseCases(repo *repo.UserRepository) *UseCases {
	return &UseCases{repo}
}

func (actions *UseCases) CreateUser(user user.User) {
	_, err := actions.repo.CreateUser(context.Background(), user)
	if err != nil {
		return
	}

}
