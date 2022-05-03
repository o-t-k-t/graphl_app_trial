package repository

import (
	"context"

	"github.com/o-t-k-t/graphl_app_trial/app/entity"
	"github.com/o-t-k-t/graphl_app_trial/ent"
	"github.com/o-t-k-t/graphl_app_trial/ent/user"
)

type UserRepository struct {
	EntClient *ent.Client
}

func (ur UserRepository) FindUser(id int) (entity.User, error) {
	u, err := ur.EntClient.User.Get(context.Background(), id)
	if err != nil {
		return entity.User{}, err
	}

	return entity.User(*u), nil
}

func (ur UserRepository) FindUsers() ([]entity.User, error) {
	us, err := ur.EntClient.User.Query().All(context.Background())
	if err != nil {
		return nil, err
	}

	return ur.toEntitySlice(us), nil
}

func (ur UserRepository) FindUsersByName(name string) ([]entity.User, error) {
	us, err := ur.EntClient.
		User.Query().Where(user.NameContains(name)).All(context.Background())
	if err != nil {
		return nil, err
	}

	return ur.toEntitySlice(us), nil
}

func (ur UserRepository) Create(u entity.User) (entity.User, error) {
	record, err := ur.EntClient.
		User.
		Create().
		SetAge(u.Age).
		SetName(u.Name).
		Save(context.Background())
	if err != nil {
		return entity.User{}, err
	}

	return entity.User(*record), nil
}

func (ur UserRepository) toEntitySlice(us []*ent.User) []entity.User {
	eus := make([]entity.User, len(us))
	for _, u := range us {
		eus = append(eus, entity.User(*u))
	}
	return eus
}
