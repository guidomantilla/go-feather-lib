package properties

import (
	"sync/atomic"
)

var singleton atomic.Value

func retrieveSingleton() Properties {
	value := singleton.Load()
	if value == nil {
		return Default()
	}
	return value.(Properties)
}

func Default() Properties {
	properties := NewDefaultProperties()
	singleton.Store(properties)
	return properties
}

func Custom(slice []string) Properties {
	properties := NewDefaultProperties(FromSlice(slice))
	singleton.Store(properties)
	return properties
}

func Add(property string, value string) {
	properties := retrieveSingleton()
	properties.Add(property, value)
}

func Get(property string) string {
	properties := retrieveSingleton()
	return properties.Get(property)
}

func AsMap() map[string]string {
	properties := retrieveSingleton()
	return properties.AsMap()
}
