package cmd

import (
	"fmt"
	"os"
	"strings"
)

func handleSingleQuotes(str string) string {
	str = strings.ReplaceAll(str, "'", "")
	return str
}
func Echo(argv []string, stderr *os.File, stdout *os.File) {
	if len(argv) == 0 {
		return
	}
	toEcho := strings.Join(argv, " ")
	toEcho = handleSingleQuotes(toEcho)

	fmt.Fprintln(stdout, toEcho)
}
