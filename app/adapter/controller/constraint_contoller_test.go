package controller

import (
	"testing"
)

func TestConstraintController_validate(t *testing.T) {
	type args struct {
		obj       interface{}
		field     *string
		minLength *int
		maxLength *int
		min       *int
		max       *int
	}
	sp := func(s string) *string { return &s }
	ip := func(i int) *int { return &i }

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"If int value in range, return success", args{map[string]interface{}{"field": int64(100)}, sp("field"), nil, nil, ip(100), ip(100)}, false},
		{"If int value falls min, return error", args{map[string]interface{}{"field": int64(100)}, sp("field"), nil, nil, ip(101), nil}, true},
		{"If int value exceeds max, return error", args{map[string]interface{}{"field": int64(100)}, sp("field"), nil, nil, nil, ip(99)}, true},
		{"If string length in range, return success", args{map[string]interface{}{"field": "hello"}, sp("field"), ip(5), ip(5), nil, nil}, false},
		{"If string length falls min, return error", args{map[string]interface{}{"field": "hello"}, sp("field"), ip(6), nil, nil, nil}, true},
		{"If string length exceeds max, return error", args{map[string]interface{}{"field": "hello"}, sp("field"), nil, ip(4), nil, nil}, true},
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
