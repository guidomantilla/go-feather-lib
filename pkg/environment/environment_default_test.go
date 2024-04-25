package environment

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/guidomantilla/go-feather-lib/pkg/properties"
)

func Test_NewDefaultEnvironment(t *testing.T) {

	source := properties.NewDefaultPropertySource("some_property_source", properties.NewDefaultProperties())
	environment := NewDefaultEnvironment(WithPropertySources(source))

	assert.NotNil(t, environment)
	assert.NotNil(t, environment.propertySources)
	assert.NotEmpty(t, environment.propertySources)
	assert.Equal(t, "some_property_source", environment.propertySources[0].AsMap()["name"])
}

func Test_GetValue(t *testing.T) {

	props := properties.NewDefaultProperties()
	props.Add("some_property", "some_value")

	source := properties.NewDefaultPropertySource("some_property_source", props)
	environment := NewDefaultEnvironment(WithPropertySources(source))

	value := environment.GetValue("some_property")
	assert.NotNil(t, value)
	assert.NotEmpty(t, value)
	assert.Equal(t, "some_value", value.AsString())
}

func Test_GetValueOrDefault_ReturnDefault(t *testing.T) {

	props := properties.NewDefaultProperties()

	source := properties.NewDefaultPropertySource("some_property_source", props)
	environment := NewDefaultEnvironment(WithPropertySources(source))

	value := environment.GetValueOrDefault("some_property", "some_default_value")
	assert.NotNil(t, value)
	assert.NotEmpty(t, value)
	assert.Equal(t, "some_default_value", value.AsString())
}

func Test_GetValueOrDefault_ReturnValue(t *testing.T) {

	props := properties.NewDefaultProperties()
	props.Add("some_property", "some_value")

	source := properties.NewDefaultPropertySource("some_property_source", props)
	environment := NewDefaultEnvironment(WithPropertySources(source))

	value := environment.GetValueOrDefault("some_property", "some_default_value")
	assert.NotNil(t, value)
	assert.NotEmpty(t, value)
	assert.Equal(t, "some_value", value.AsString())
}

func Test_GetPropertySources(t *testing.T) {

	props := properties.NewDefaultProperties()
	props.Add("some_property", "some_value")

	source := properties.NewDefaultPropertySource("some_property_source", props)
	environment := NewDefaultEnvironment(WithPropertySources(source))

	propertySources := environment.GetPropertySources()
	assert.NotNil(t, propertySources)
	assert.NotEmpty(t, propertySources)
	assert.Equal(t, "some_property_source", environment.propertySources[0].AsMap()["name"])
}

func Test_AppendPropertySources(t *testing.T) {

	props1 := properties.NewDefaultProperties()
	props1.Add("some_property", "some_value")

	source1 := properties.NewDefaultPropertySource("some_property_source1", props1)
	environment := NewDefaultEnvironment(WithPropertySources(source1))

	props2 := properties.NewDefaultProperties()
	source2 := properties.NewDefaultPropertySource("some_property_source2", props2)

	environment.AppendPropertySources(source2)

	propertySources := environment.GetPropertySources()

	assert.NotNil(t, propertySources)
	assert.NotEmpty(t, propertySources)
	assert.Equal(t, "some_property_source1", environment.propertySources[0].AsMap()["name"])
	assert.Equal(t, "some_property_source2", environment.propertySources[1].AsMap()["name"])
}

func Test_WithArraySource(t *testing.T) {

	environment := NewDefaultEnvironment(WithArraySource("some_property_source3", []string{"some_property", "some_value"}))

	propertySources := WithArraySource("some_property_source3", []string{"some_property", "some_value"})

	assert.NotNil(t, propertySources)
	assert.NotEmpty(t, propertySources)
	assert.Equal(t, "some_property_source3", environment.propertySources[0].AsMap()["name"])
}

func Test_WithArrays(t *testing.T) {

	osArgs := os.Environ()
	cmdArgs := []string{"some_property", "some_value"}
	environment := NewDefaultEnvironment(WithArrays(osArgs, cmdArgs))

	assert.Equal(t, OsPropertySourceName, environment.propertySources[0].AsMap()["name"])
	assert.Equal(t, CmdPropertySourceName, environment.propertySources[1].AsMap()["name"])
}
