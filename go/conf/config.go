package conf

type Config interface {
	Port() int
	ViewFileHome() string
	ViewFileSuffix() string
}

type defaultConfig struct {
	port int
	viewFileHome string
	viewFileSuffix string
}

func NewConfig() Config {
	ret := defaultConfig{port: 1209, viewFileHome: "templates", viewFileSuffix: ".html"}
	return ret
}

func (c defaultConfig) Port() int {
	return c.port
}

func (c defaultConfig) ViewFileHome() string {
	return c.viewFileHome
}

func (c defaultConfig) ViewFileSuffix() string {
	return c.viewFileSuffix
}
