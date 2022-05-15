package controller

import (
	"context"
	"fmt"
	"log"

	"github.com/99designs/gqlgen/graphql"
)

type ConstraintController struct{}

func NewConstraintController() ConstraintController {
	return ConstraintController{}
}

func (c ConstraintController) Constraint(
	ctx context.Context,
	obj interface{},
	next graphql.Resolver,
	minLength *int,
	maxLength *int,
	min *int,
	max *int,
) (
	res interface{},
	err error,
) {
	field := graphql.GetPathContext(ctx).Field

	if err := c.validate(obj, field, minLength, maxLength, min, max); err != nil {
		log.Println("AAAAAAAAAAABBBBBBBBBBBBB")
		return nil, err
	}

	return next(ctx)
}

func (c ConstraintController) validate(
	obj interface{},
	field *string,
	minLength *int,
	maxLength *int,
	min *int,
	max *int,
) (
	err error,
) {
	objMap, ok := obj.(map[string]interface{})
	if !ok {
		return fmt.Errorf("illegal obj. %+v", obj)
	}

	val, ok := objMap[*field]
	if !ok {
		return fmt.Errorf("field not found. %+v, %s", objMap, *field)
	}

	// gqlgen's validation checks that flied type and presence matche with the schema. so we don't. (focus only value)
	// for String
	strVal, ok := val.(string)
	if ok {
		len := len(strVal)
		if minLength != nil {
			if len < *minLength {
				return fmt.Errorf("field too short. %s, %d", *field, len)
			}
		}

		if maxLength != nil {
			if len > *maxLength {
				return fmt.Errorf("field too long. %s, %d", *field, len)
			}
		}
	}

	// for Int
	intVal, ok := val.(int64)
	if ok {
		if min != nil {
			if intVal < int64(*min) {
				return fmt.Errorf("field too small. %s, %d", *field, intVal)
			}
		}

		if max != nil {
			if intVal > int64(*max) {
				return fmt.Errorf("field too large. %s, %d", *field, intVal)
			}
		}
	}

	return nil
}
