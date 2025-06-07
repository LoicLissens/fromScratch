package cmd

import (
	"fmt"
	"os"
)

func Pwd(argv []string, stderr *os.File, stdout *os.File) {
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(stderr, "Error getting current directory:", err)
		return
	}
	fmt.Fprintln(stdout, currentDir)
}
