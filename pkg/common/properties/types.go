package properties

var (
	_ Properties     = (*DefaultProperties)(nil)
	_ PropertySource = (*DefaultPropertySource)(nil)
	_ PropertySource = (*MockPropertySource)(nil)
	_ Properties     = (*MockProperties)(nil)
)

type Properties interface {
	Add(property string, value string)
	Get(property string) string
	AsMap() map[string]string
}

type PropertySource interface {
	Get(property string) string
	AsMap() map[string]any
}
