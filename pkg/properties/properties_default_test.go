package properties

import (
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_PropertiesAdd(t *testing.T) {

	properties := &DefaultProperties{
		internalMap: make(map[string]string),
	}
	properties.Add("property01", "value01")
	properties.Add("property02", "value02")
	properties.Add("property03", "value03")

	assert.Equal(t, "value01", properties.internalMap["property01"])
	assert.Equal(t, "value02", properties.internalMap["property02"])
	assert.Equal(t, "value03", properties.internalMap["property03"])

}

func Test_PropertiesGet(t *testing.T) {

	properties := &DefaultProperties{
		internalMap: make(map[string]string),
	}
	properties.Add("property01", "value01")
	properties.Add("property02", "value02")
	properties.Add("property03", "value03")

	assert.Equal(t, "value01", properties.Get("property01"))
	assert.Equal(t, "value02", properties.Get("property02"))
	assert.Equal(t, "value03", properties.Get("property03"))
}

func Test_PropertiesAsMap(t *testing.T) {

	properties := &DefaultProperties{
		internalMap: make(map[string]string),
	}
	properties.Add("property01", "value01")
	properties.Add("property02", "value02")
	properties.Add("property03", "value03")

	internalMap := properties.AsMap()

	assert.Equal(t, properties.internalMap, internalMap)
}

func Test_NewDefaultProperties(t *testing.T) {
	properties := NewDefaultProperties()

	assert.NotNil(t, properties)
	assert.NotNil(t, properties.internalMap)
}

func Test_PropertiesBuilderFromArray_Ok(t *testing.T) {

	osArgs := os.Environ()
	properties := NewDefaultProperties(FromArray(osArgs))

	assert.NotNil(t, properties)
	assert.Equal(t, len(osArgs), len(properties.internalMap))
}

func Test_PropertiesBuilderFromArray_Error(t *testing.T) {

	args := []string{"ola", "adios"}
	properties := NewDefaultProperties(FromArray(args))

	assert.NotNil(t, properties)
	assert.Equal(t, 0, len(properties.internalMap))
}

func TestFromArray(t *testing.T) {
	type args struct {
		array []string
	}
	tests := []struct {
		name string
		args args
		want DefaultPropertiesOption
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromArray(tt.args.array); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewDefaultProperties(t *testing.T) {
	type args struct {
		options []DefaultPropertiesOption
	}
	tests := []struct {
		name string
		args args
		want *DefaultProperties
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDefaultProperties(tt.args.options...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDefaultProperties() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDefaultProperties_Add(t *testing.T) {
	type args struct {
		property string
		value    string
	}
	tests := []struct {
		name string
		p    *DefaultProperties
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.p.Add(tt.args.property, tt.args.value)
		})
	}
}

func TestDefaultProperties_Get(t *testing.T) {
	type args struct {
		property string
	}
	tests := []struct {
		name string
		p    *DefaultProperties
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Get(tt.args.property); got != tt.want {
				t.Errorf("DefaultProperties.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDefaultProperties_AsMap(t *testing.T) {
	tests := []struct {
		name string
		p    *DefaultProperties
		want map[string]string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.AsMap(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DefaultProperties.AsMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
