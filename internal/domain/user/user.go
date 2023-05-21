package user

import (
	"errors"
	"fmt"
)

type User struct {
	id   *Id
	name string
}

func NewUser(name string) *User {

	// idの設定は、ストレージの保存時に行う
	// そのため、アプリケーション内のUserの新規作成時にはidは設定しない

	user := new(User)
	user.name = name

	return user
}

func ReconstructUser(id int64, name string) (*User, error) {

	userId, err := NewId(id)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("failed to reconstruct the user :%v", err))
	}

	user := new(User)
	user.id = userId
	user.name = name

	return user, nil
}

func (u *User) Id() *Id {
	return u.id
}

func (u *User) Name() string {
	return u.name
}
