package application

import (
	"log"
	"strings"
)

func ShowHelp() {
	log.Println("Usage: checkVersionJetBrains <path to product-info.yaml>")
	log.Println("Example: checkVersionJetBrains ~/jetBrains_config.yaml")
	log.Println(strings.Repeat("-", 75))
	log.Println("The file must follow the following structure:")
	log.Println("BasePath: ~/bin")
	log.Println("Products:")
	log.Println("  - Name: GoLand")
	log.Println("    Path: bin/Go/GoLand")
	log.Println("  - Name: WebStorm")
	log.Println("    Path: bin/WebStorm/WebStorm")
}
