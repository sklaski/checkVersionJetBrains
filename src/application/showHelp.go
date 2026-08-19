package application

import (
	"log"
	"strings"

	"checkVersionJetBrains/src/domain"
)

func ShowHelp() {
	log.Println("Usage: checkVersionJetBrains <path to product-info.yaml>")
	log.Println("Example: checkVersionJetBrains ~/jetBrains_config.yaml")
	log.Println(strings.Repeat("-", 75))
	log.Println("The file must follow the following structure:")
	log.Println("basePath: ~/bin")
	log.Println("products:")
	log.Println("  - name: GoLand")
	log.Println("    path: bin/Go/GoLand")
	log.Println("  - name: WebStorm")
	log.Println("    path: bin/WebStorm/WebStorm")
	log.Println("------------------")
	log.Println("DefaultPaths:")
	for _, path := range domain.DefaultPaths {
		log.Println(path)
	}
}
