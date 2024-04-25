package properties

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewDefaultPropertySource(t *testing.T) {

	properties := &DefaultProperties{
		internalMap: make(map[string]string),
	}
	properties.Add("property01", "value01")
	properties.Add("property02", "value02")
	properties.Add("property03", "value03")

	propertySource := NewDefaultPropertySource("group", properties)

	assert.NotNil(t, propertySource)
	assert.Equal(t, properties, propertySource.properties)
	assert.Equal(t, "group", propertySource.name)
}

func Test_PropertySourceGet(t *testing.T) {

	properties := &DefaultProperties{
		internalMap: make(map[string]string),
	}
	properties.Add("property01", "value01")
	properties.Add("property02", "value02")
	properties.Add("property03", "value03")

	propertySource := NewDefaultPropertySource("group", properties)

	value := propertySource.Get("property01")

	assert.Equal(t, "value01", value)
}

func Test_PropertySourceAsMap(t *testing.T) {

	properties := &DefaultProperties{
		internalMap: make(map[string]string),
	}
	properties.Add("property01", "value01")
	properties.Add("property02", "value02")
	properties.Add("property03", "value03")

	propertySource := NewDefaultPropertySource("group", properties)

	internalMap := propertySource.AsMap()

	assert.Equal(t, "group", internalMap["name"])
	assert.Equal(t, properties.internalMap, internalMap["value"])
}

func TestDefaultPropertySource_Get(t *testing.T) {
	type args struct {
		property string
	}
	tests := []struct {
		name   string
		source *DefaultPropertySource
		args   args
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.source.Get(tt.args.property); got != tt.want {
				t.Errorf("DefaultPropertySource.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDefaultPropertySource_AsMap(t *testing.T) {
	tests := []struct {
		name   string
		source *DefaultPropertySource
		want   map[string]any
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.source.AsMap(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DefaultPropertySource.AsMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewDefaultPropertySource(t *testing.T) {
	type args struct {
		name       string
		properties Properties
	}
	tests := []struct {
		name string
		args args
		want *DefaultPropertySource
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDefaultPropertySource(tt.args.name, tt.args.properties); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDefaultPropertySource() = %v, want %v", got, tt.want)
			}
		})
	}
}
