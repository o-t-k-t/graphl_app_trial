package entity

import "github.com/o-t-k-t/graphl_app_trial/ent"

type User ent.User

type UserCars struct {
	UserID int
	Count  int
}
