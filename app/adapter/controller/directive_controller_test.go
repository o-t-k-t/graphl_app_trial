package controller

import (
	"context"
	"reflect"
	"testing"

	"github.com/99designs/gqlgen/graphql"
)

func TestNewConstraintController(t *testing.T) {
	tests := []struct {
		name string
		want ConstraintController
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewConstraintController(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewConstraintController() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConstraintController_Constraint(t *testing.T) {
	type args struct {
		ctx       context.Context
		obj       interface{}
		next      graphql.Resolver
		minLength *int
		maxLength *int
		min       *int
		max       *int
	}
	tests := []struct {
		name    string
		c       ConstraintController
		args    args
		wantRes interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := ConstraintController{}
			gotRes, err := c.Constraint(tt.args.ctx, tt.args.obj, tt.args.next, tt.args.minLength, tt.args.maxLength, tt.args.min, tt.args.max)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConstraintController.Constraint() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("ConstraintController.Constraint() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func TestConstraintController_validate(t *testing.T) {
	type args struct {
		obj       interface{}
		field     *string
		minLength *int
		maxLength *int
		min       *int
		max       *int
	}
	tests := []struct {
		name    string
		c       ConstraintController
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := ConstraintController{}
			if err := c.validate(tt.args.obj, tt.args.field, tt.args.minLength, tt.args.maxLength, tt.args.min, tt.args.max); (err != nil) != tt.wantErr {
				t.Errorf("ConstraintController.validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
