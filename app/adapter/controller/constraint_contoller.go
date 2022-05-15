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
	if err := c.validateFieldIfString(val, minLength, maxLength); err != nil {
		return err
	}

	if err := c.validateFieldIfInt(val, min, max); err != nil {
		return err
	}

	return nil
}

func (c ConstraintController) validateFieldIfString(val interface{}, minLength, maxLength *int) error {
	strVal, ok := val.(string)
	if !ok {
		return nil
	}

	len := len(strVal)

	if minLength != nil {
		if len < *minLength {
			return fmt.Errorf("field too short. %d", len)
		}
	}

	if maxLength != nil {
		if len > *maxLength {
			return fmt.Errorf("field too long. %d", len)
		}
	}
	return nil
}

func (c ConstraintController) validateFieldIfInt(val interface{}, min, max *int) error {
	intVal, ok := val.(int64)
	if !ok {
		return nil
	}

	if min != nil {
		if intVal < int64(*min) {
			return fmt.Errorf("field too small. %d", intVal)
		}
	}

	if max != nil {
		if intVal > int64(*max) {
			return fmt.Errorf("field too large. %d", intVal)
		}
	}
	return nil
}
