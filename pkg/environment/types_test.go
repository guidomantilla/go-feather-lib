package environment

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewEnvVar(t *testing.T) {

	aux := "some string"
	evnVar := NewEnvVar(aux)
	assert.Equal(t, aux, evnVar.AsString())
}

func Test_EnvVarAsString(t *testing.T) {

	aux := "some string"
	evnVar := NewEnvVar(aux)
	assert.Equal(t, aux, evnVar.AsString())
}

func Test_EnvVarAsInt_Ok(t *testing.T) {

	aux := "101010"
	evnVar := NewEnvVar(aux)
	value, _ := evnVar.AsInt()
	assert.Equal(t, 101010, value)
}

func Test_EnvVarAsInt_Error(t *testing.T) {

	aux := "some string"
	evnVar := NewEnvVar(aux)
	value, _ := evnVar.AsInt()
	assert.Equal(t, 0, value)
}

func TestNewEnvVar(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want EnvVar
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewEnvVar(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEnvVar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEnvVar_AsInt(t *testing.T) {
	tests := []struct {
		name    string
		envVar  EnvVar
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.envVar.AsInt()
			if (err != nil) != tt.wantErr {
				t.Errorf("EnvVar.AsInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("EnvVar.AsInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEnvVar_AsString(t *testing.T) {
	tests := []struct {
		name   string
		envVar EnvVar
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.envVar.AsString(); got != tt.want {
				t.Errorf("EnvVar.AsString() = %v, want %v", got, tt.want)
			}
		})
	}
}
