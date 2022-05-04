package model

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// for individual field
type UserCars struct {
	UserID int `json:"userId"`
	Count  int `json:"count"`
}
