package conf

var port int = 1209
var publicDirHome string = "templates"
var viewFileSuffix string = ".html"

func Port() int {
	return port
}

func PublicDirHome() string {
	return publicDirHome
}

func ViewFileSuffix() string {
	return viewFileSuffix
}
