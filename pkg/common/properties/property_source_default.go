package properties

// DefaultPropertySource

type DefaultPropertySource struct {
	name        string
	properties  Properties
	internalMap map[string]any
}

func (source *DefaultPropertySource) Get(property string) string {
	return source.properties.Get(property)
}

func (source *DefaultPropertySource) AsMap() map[string]any {
	return source.internalMap
}

//

func NewDefaultPropertySource(name string, properties Properties) *DefaultPropertySource {

	internalMap := make(map[string]any)
	internalMap["name"], internalMap["value"] = name, properties.AsMap()

	return &DefaultPropertySource{
		name:        name,
		properties:  properties,
		internalMap: internalMap,
	}
}
