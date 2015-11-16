package conf

type Config interface {
	Port() int
	PublicDirHome() string
	ViewFileSuffix() string
}

type defaultConfig struct {
	port int
	publicDirHome string
	viewFileSuffix string
}

func NewConfig() Config {
	ret := defaultConfig{port: 1209, publicDirHome: "templates", viewFileSuffix: ".html"}
	return ret
}

func (c defaultConfig) Port() int {
	return c.port
}

func (c defaultConfig) PublicDirHome() string {
	return c.publicDirHome
}

func (c defaultConfig) ViewFileSuffix() string {
	return c.viewFileSuffix
}
