package controller

import (
	"context"
	"fmt"

	"github.com/graph-gophers/dataloader"
	"github.com/o-t-k-t/graphl_app_trial/app/adapter/repository"
	"github.com/o-t-k-t/graphl_app_trial/app/entity"
	"github.com/o-t-k-t/graphl_app_trial/app/usecase"
	"github.com/o-t-k-t/graphl_app_trial/ent"
	"github.com/o-t-k-t/graphl_app_trial/graph/model"
)

// UserController passes the data passed from the client to app. logic.
// Then, converts data returned from app logic to a view model for the client and passes.
// A controller has action method for entrypoint. such as resolver, dataloader, batch join etc.
type UserController struct {
	UserUsecase usecase.UserUsecase
}

func NewUserController(entClient *ent.Client) UserController {
	return UserController{
		UserUsecase: usecase.UserUsecase{
			UserRepository: repository.UserRepository{
				EntClient: entClient,
			},
		},
	}
}

func (c UserController) Create(ctx context.Context, input model.UserInput) (*model.User, error) {
	// translate input into domain entity.
	user, err := entity.NewUser(input.Name, input.Age)
	if err != nil {
		return nil, fmt.Errorf("create failed %w", err)
	}

	// execute application logic.
	createdUser, err := c.UserUsecase.CreateUser(user)
	if err != nil {
		return nil, fmt.Errorf("create failed %w", err)
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

func (c UserController) BatchFindCars(ctx context.Context, keys DataloaderKeys) []*dataloader.Result {
	// translate input into domain entity.
	ids, _ := keys.ToIDsAndKeyOrders()

	// execute application logic.
	counts, err := c.UserUsecase.BatchFindCars(ids)
	if err != nil {
		return []*dataloader.Result{{Data: nil, Error: err}}
	}

	// translate domain entity into view model
	results := keys.ToResult(func(id int) (interface{}, error) {
		for _, c := range counts {
			if c.UserID == id {
				return c, nil
			}
		}
		return nil, fmt.Errorf("not found %d", id)
	})

	return results
}

func toViewUser(u entity.User) *model.User {
	return &model.User{
		ID:   fmt.Sprintf("%d", u.ID),
		Name: u.Name,
	}
}
