package user

import "errors"

type Id struct {
	value int64
}

func NewId(value int64) (*Id, error) {
	if value <= 0 {
		return nil, errors.New("the ID must be a value greater than or equal to 1")
	}

	Id := new(Id)
	Id.value = value

	return Id, nil
}

func (id *Id) Value() int64 {
	return id.value
}
