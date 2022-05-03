package controller

import (
	"context"
	"fmt"

	"github.com/o-t-k-t/graphl_app_trial/app/entity"
	"github.com/o-t-k-t/graphl_app_trial/app/usecase"
	"github.com/o-t-k-t/graphl_app_trial/graph/model"
)

// UserController passes the data passed from the client to app. logic.
// Then, converts data returned from app logic to a view model for the client and passes.
type UserController struct {
	UserUsecase usecase.UserUsecase
}

func (c UserController) Create(ctx context.Context, input model.UserInput) (*model.User, error) {
	// translate input into domain entity.
	userInput := entity.User{
		Name: input.Name,
		Age:  input.Age,
	}

	// execute application logic.
	createdUser, err := c.UserUsecase.CreateUser(userInput)
	if err != nil {
		return nil, err
	}

	// translate domain entity into view model
	return toViewUser(createdUser), nil
}

func (c UserController) List(ctx context.Context) ([]*model.User, error) {
	// execute application logic.
	entityUsers, err := c.UserUsecase.FindUsers()
	if err != nil {
		return nil, err
	}

	// translate domain entity into view model
	var viewUsers []*model.User
	for _, u := range entityUsers {
		viewUsers = append(viewUsers, toViewUser(u))
	}
	return viewUsers, nil
}

func toViewUser(u entity.User) *model.User {
	return &model.User{
		ID:   fmt.Sprintf("%d", u.ID),
		Name: u.Name,
	}
}
