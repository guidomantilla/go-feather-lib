package properties

import (
	"reflect"
	"testing"
)

func Test_retrieveSingleton(t *testing.T) {
	tests := []struct {
		name string
		want Properties
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := retrieveSingleton(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("retrieveSingleton() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDefault(t *testing.T) {
	tests := []struct {
		name string
		want Properties
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Default(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Default() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCustom(t *testing.T) {
	type args struct {
		array []string
	}
	tests := []struct {
		name string
		args args
		want Properties
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Custom(tt.args.array); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Custom() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAdd(t *testing.T) {
	type args struct {
		property string
		value    string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Add(tt.args.property, tt.args.value)
		})
	}
}

func TestGet(t *testing.T) {
	type args struct {
		property string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Get(tt.args.property); got != tt.want {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAsMap(t *testing.T) {
	tests := []struct {
		name string
		want map[string]string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AsMap(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AsMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
