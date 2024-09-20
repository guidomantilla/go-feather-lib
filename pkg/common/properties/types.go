package properties

var (
	_ Properties       = (*properties)(nil)
	_ PropertiesSource = (*propertiesSource)(nil)
	_ PropertiesSource = (*MockPropertiesSource)(nil)
	_ Properties       = (*MockProperties)(nil)
)

type PropertiesOption func(properties *properties)

type Properties interface {
	Add(property string, value string)
	Get(property string) string
	AsMap() map[string]string
}

//

type PropertiesSource interface {
	Get(property string) string
	AsMap() map[string]any
}
