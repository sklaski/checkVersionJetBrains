package application

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

func FullQualifiedPath(arg string) string {
	if strings.HasPrefix(arg, "~/") {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			log.Fatal(err)
		}
		arg = filepath.Join(homeDir, arg[2:])
	}
	return arg
}
