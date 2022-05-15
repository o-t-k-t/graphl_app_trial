package entity

import (
	"errors"

	"github.com/o-t-k-t/graphl_app_trial/ent"
)

type User ent.User

type UserCars struct {
	UserID int
	Count  int
}

func NewUser(name string, age int) (User, error) {
	if len(name) >= 200 {
		return User{}, errors.New("name length too long.")
	}

	if age >= 200 {
		return User{}, errors.New("age too large.")
	}

	return User{Name: name, Age: age}, nil
}
