package properties

import (
	"sync/atomic"
)

var singleton atomic.Value

func instance() Properties {
	value := singleton.Load()
	if value == nil {
		return Load()
	}
	return value.(Properties)
}

func Load(slices ...[]string) Properties {
	withSlices := make([]Option, 0)
	for _, slice := range slices {
		withSlices = append(withSlices, FromSlice(slice))
	}
	properties := New(withSlices...)
	singleton.Store(properties)
	return properties
}

func Add(property string, value string) {
	instance().Add(property, value)
}

func Get(property string) string {
	return instance().Get(property)
}

func AsMap() map[string]string {
	return instance().AsMap()
}
