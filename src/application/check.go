package application

import (
	"fmt"
	"log"
	"os"

	"checkVersionJetBrains/src/domain"
)

func CheckParam(args []string) string {
	if len(args) > 2 {
		ShowHelp()
		os.Exit(1)
	}
	var arg string
	if len(args) == 2 {
		if args[1] == "--help" {
			ShowHelp()
			os.Exit(0)
		}
		arg = FullQualifiedPath(os.Args[1])
	}

	if len(arg) == 0 {
		var err error
		arg, err = CheckDefaultPaths()
		if err != nil {
			log.Fatal(err)
		}
	}

	if _, err := os.Stat(arg); os.IsNotExist(err) {
		log.Fatalf("File %s does not exist", arg)
	}
	return arg
}

func CheckDefaultPaths() (string, error) {
	for _, path := range domain.DefaultPaths {
		fullQualifiedPath := FullQualifiedPath(path)
		if _, err := os.Stat(fullQualifiedPath); err == nil {
			return fullQualifiedPath, nil
		}
	}
	return "", fmt.Errorf("no configuration file found in default paths: %v", domain.DefaultPaths)
}
