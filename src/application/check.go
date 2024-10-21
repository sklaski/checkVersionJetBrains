package application

import (
	"log"
	"os"
)

func CheckParam(args []string) string {
	if len(args) != 2 || args[1] == "--help" {
		ShowHelp()
		os.Exit(1)
	}
	arg := FullQualifiedPath(os.Args[1])
	if _, err := os.Stat(arg); os.IsNotExist(err) {
		log.Fatalf("File %s does not exist", arg)
	}
	return arg
}
