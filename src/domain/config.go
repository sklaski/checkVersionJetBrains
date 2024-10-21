package domain

var DefaultPaths = []string{
	"~/jetbrains.yaml",
	"/etc/jetbrains.yaml",
	"~/.config/jetbrains.yaml",
}

type (
	ConfigFile struct {
		BasePath       string         `yaml:"BasePath"`
		ProductConfigs ProductConfigs `yaml:"Products"`
	}

	ProductConfigs []ProductConfig
	ProductConfig  struct {
		Name string `yaml:"Name"`
		Path string `yaml:"Path"`
	}
)
