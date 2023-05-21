package user

import "context"

type Repository interface {
	FindById(ctx context.Context, userId *Id) (*User, error)
	Save(ctx context.Context, user *User) error
}
