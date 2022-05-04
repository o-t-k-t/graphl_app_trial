package controller

import (
	"github.com/graph-gophers/dataloader"
)

type DataloaderKeys interface {
	ToIDsAndKeyOrders() ([]int, map[int]int)
	ToResult(f func(id int) (interface{}, error)) []*dataloader.Result
}
