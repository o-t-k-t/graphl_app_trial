package usecase

import (
	"github.com/o-t-k-t/graphl_app_trial/app/entity"
)

type UserRepository interface {
	FindUsers() ([]entity.User, error)
	Create(u entity.User) (entity.User, error)
}

type UserUsecase struct {
	UserRepository UserRepository
}

func (uc UserUsecase) FindUsers() ([]entity.User, error) {
	return uc.UserRepository.FindUsers()
}

func (uc UserUsecase) CreateUser(u entity.User) (entity.User, error) {
	return uc.UserRepository.Create(u)
}
