package usecase

import (
	"context"
	"errors"
	"fmt"
	"generate-id-use-autoincrement/internal/domain/user"
)

type CreateUserUsecase interface {
	Execute(ctx context.Context, name string) error
}

type createUserUsecase struct {
	userRepository user.Repository
}

func NewCreateUserUsecase(repository user.Repository) CreateUserUsecase {
	return &createUserUsecase{userRepository: repository}
}

func (u *createUserUsecase) Execute(ctx context.Context, name string) error {

	// ユーザーの新規作成
	// この時にユーザーのIDは設定しない
	newUser := user.NewUser(name)

	// ユーザーの保存
	// 保存時に初めて、IDが設定される
	err := u.userRepository.Save(ctx, newUser)
	if err != nil {
		return errors.New(fmt.Sprintf("fail to create user: %v", err))
	}

	return nil
}
