package cmd

import (
	"os"
)

func Exit(argv []string, stderr *os.File, stdout *os.File) {
	if len(argv) == 0 {
		return
	}
	if argv[0] == "0" {
		os.Exit(0)
	}
}
