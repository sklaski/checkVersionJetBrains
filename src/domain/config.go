package domain

var DefaultPaths = []string{
	"~/jetbrains.yaml",
	"/etc/jetbrains.yaml",
	"~/.config/jetbrains.yaml",
}

type (
	ConfigFile struct {
		BasePath       string         `yaml:"basePath"`
		ProductConfigs ProductConfigs `yaml:"products"`
	}

	ProductConfigs []ProductConfig
	ProductConfig  struct {
		Name string `yaml:"name"`
		Path string `yaml:"path"`
	}
)
