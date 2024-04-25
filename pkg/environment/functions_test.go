package environment

import (
	"os"
	"reflect"
	"testing"

	"github.com/guidomantilla/go-feather-lib/pkg/properties"
)

func Test_retrieveSingleton(t *testing.T) {

	envs := os.Environ()
	env := NewDefaultEnvironment(WithArraySource(OsPropertySourceName, envs))

	tests := []struct {
		name string
		want Environment
	}{
		{
			name: "Value is nil Path",
			want: env,
		},
		{
			name: "Value is not nil  Path",
			want: env,
		},
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

	envs := os.Environ()
	env := NewDefaultEnvironment(WithArraySource(OsPropertySourceName, envs))

	tests := []struct {
		name string
		want Environment
	}{
		{
			name: "Happy Path",
			want: env,
		},
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

	osArgs := os.Environ()
	cmdArgs := []string{"some_property=some_value"}
	env := NewDefaultEnvironment(WithArrays(osArgs, cmdArgs))

	type args struct {
		cmdArgsArray []string
	}
	tests := []struct {
		name string
		args args
		want Environment
	}{
		{
			name: "Happy Path",
			args: args{
				cmdArgsArray: cmdArgs,
			},
			want: env,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Custom(tt.args.cmdArgsArray); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Custom() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetValue(t *testing.T) {
	cmdArgs := []string{"some_property=some_value"}
	Custom(cmdArgs)
	type args struct {
		property string
	}
	tests := []struct {
		name string
		args args
		want EnvVar
	}{
		{
			name: "Happy Path",
			args: args{
				property: "some_property",
			},
			want: "some_value",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetValue(tt.args.property); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetValueOrDefault(t *testing.T) {
	cmdArgs := []string{"some_property=some_value"}
	Custom(cmdArgs)
	type args struct {
		property     string
		defaultValue string
	}
	tests := []struct {
		name string
		args args
		want EnvVar
	}{
		{
			name: "Happy Path",
			args: args{
				property:     "some_property",
				defaultValue: "some_value2",
			},
			want: "some_value",
		},
		{
			name: "Other Path",
			args: args{
				property:     "some_property2",
				defaultValue: "some_value2",
			},
			want: "some_value2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetValueOrDefault(tt.args.property, tt.args.defaultValue); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetValueOrDefault() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetPropertySources(t *testing.T) {

	env := Default()

	tests := []struct {
		name string
		want []properties.PropertySource
	}{
		{
			name: "Happy Path",
			want: env.GetPropertySources(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetPropertySources(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPropertySources() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAppendPropertySources(t *testing.T) {
	type args struct {
		propertySources []properties.PropertySource
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "",
			args: args{
				propertySources: []properties.PropertySource{
					properties.NewDefaultPropertySource("some_property_source", properties.NewDefaultProperties()),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AppendPropertySources(tt.args.propertySources...)
		})
	}
}
