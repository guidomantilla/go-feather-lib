package properties

import (
	"fmt"

	"github.com/guidomantilla/go-feather-lib/pkg/common/assert"
)

type propertiesSource struct {
	name       string
	properties Properties
	internal   map[string]any
}

func newPropertiesSource(name string, properties Properties) *propertiesSource {
	assert.NotEmpty(name, fmt.Sprintf("common properties: %s error - name is required", name))
	assert.NotNil(properties, fmt.Sprintf("common properties: %s error - handler is required", name))

	internalMap := make(map[string]any)
	internalMap["name"], internalMap["value"] = name, properties.AsMap()

	return &propertiesSource{
		name:       name,
		properties: properties,
		internal:   internalMap,
	}
}

func (source *propertiesSource) Get(property string) string {
	return source.properties.Get(property)
}

func (source *propertiesSource) AsMap() map[string]any {
	return source.internal
}
