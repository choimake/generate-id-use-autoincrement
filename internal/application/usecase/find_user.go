package usecase

import (
	"context"
	"errors"
	"fmt"
	user2 "generate-id-use-autoincrement/internal/domain/user"
)

type FindUserUsecase interface {
	Execute(ctx context.Context, id int64) (*user2.User, error)
}

type findUserUsecase struct {
	userRepository user2.Repository
}

func NewFindUserUsecase(repository user2.Repository) FindUserUsecase {
	return &findUserUsecase{userRepository: repository}
}

func (u *findUserUsecase) Execute(ctx context.Context, id int64) (*user2.User, error) {

	userId, err := user2.NewId(id)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("fail to find user: %v", err))
	}

	foundUser, err := u.userRepository.FindById(ctx, userId)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("fail to find user: %v", err))
	}

	return foundUser, nil
}
