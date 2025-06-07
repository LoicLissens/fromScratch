package utils

import (
	"os"
	"strings"
)

func FindPath(cmd string) (string, bool) {
	pathString := os.Getenv("PATH")
	for _, path := range strings.Split(pathString, ":") {
		cmdName := path + "/" + cmd
		if _, err := os.Stat(cmdName); err == nil {
			return path + "/" + cmd, true
		}
	}
	return "", false
}
