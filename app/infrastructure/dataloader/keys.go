package dataloader

import (
	"fmt"
	"strconv"

	"github.com/graph-gophers/dataloader"
)

// DataloaderKeys has translation processes relating dataloader.Keys and each app. logics.
// Since these process are frequently needed, so it is cut out.
type DataloaderKeys dataloader.Keys

// toIdsAndKeyOrders translate dataloader.Keys into ids for usecase.
func (keys DataloaderKeys) ToIDsAndKeyOrders() ([]int, map[int]int) {
	// prepare id and index pair infomation
	ids := make([]int, 0, len(keys))
	keyIndexMap := make(map[int]int, len(keys))
	for idx, key := range keys {
		userId, err := strconv.Atoi(key.String())
		if err != nil {
			continue
		}

		ids = append(ids, userId)
		keyIndexMap[userId] = idx
	}
	return ids, keyIndexMap
}

func (keys DataloaderKeys) ToResult(f func(id int) (interface{}, error)) []*dataloader.Result {
	results := make([]*dataloader.Result, len(keys))

	for i, key := range keys {
		id, err := strconv.Atoi(key.String())
		if err != nil {
			results[i] = &dataloader.Result{Data: nil, Error: fmt.Errorf("cannot convet to id %s", key)}
			continue
		}

		data, err := f(id)
		results[i] = &dataloader.Result{Data: data, Error: fmt.Errorf("cannot convet to id %s", key)}
	}

	return nil
}
