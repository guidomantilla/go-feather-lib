package properties

import (
	"sync/atomic"
)

var singleton atomic.Value

func retrieve() Properties {
	value := singleton.Load()
	if value == nil {
		return Load()
	}
	return value.(Properties)
}

func Load(slices ...[]string) Properties {
	withSlices := make([]DefaultPropertiesOption, 0)
	for _, slice := range slices {
		withSlices = append(withSlices, FromSlice(slice))
	}
	properties := NewDefaultProperties(withSlices...)
	singleton.Store(properties)
	return properties
}

func Add(property string, value string) {
	properties := retrieve()
	properties.Add(property, value)
}

func Get(property string) string {
	properties := retrieve()
	return properties.Get(property)
}

func AsMap() map[string]string {
	properties := retrieve()
	return properties.AsMap()
}
