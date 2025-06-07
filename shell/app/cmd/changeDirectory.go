package cmd

import (
	"fmt"
	"os"
	"runtime"
)

const HOMEDIR = "~"
const HOMEVARENVWINDOWS = "USERPROFILE"
const HOMEVARENV = "HOME"

func ChangeDirectory(argv []string, stderr *os.File, stdout *os.File) {
	path := argv[0]
	var pathPrefix string
	if string(path[0]) == HOMEDIR {
		var homeDir string
		if runtime.GOOS == "windows" {
			homeDir = os.Getenv(HOMEVARENVWINDOWS)
		} else {
			homeDir = os.Getenv(HOMEVARENV)
		}
		pathPrefix = homeDir
		path = path[1:]
	}
	if e := os.Chdir(pathPrefix + path); e != nil {
		fmt.Fprintln(stderr, "cd: "+argv[0]+": No such file or directory")
	}
}
