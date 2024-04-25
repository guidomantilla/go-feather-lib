package util

import (
	"reflect"
	"testing"
)

func TestValueToPtr(t *testing.T) {
	type args struct {
		value any
	}
	tests := []struct {
		name string
		args args
		want *any
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValueToPtr(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ValueToPtr() = %v, want %v", got, tt.want)
			}
		})
	}
}
