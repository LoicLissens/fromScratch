package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/app/cmd"
	"github.com/codecrafters-io/shell-starter-go/app/readline"
)

func main() {
	cmdExecuter := cmd.New(os.Stderr, os.Stdout, os.Stdin)
	rl := readline.New(readline.TTY, readline.CBREAK) // Cbreak for now (I don't have to handle CTRL + C and so on ...)
	rl.Completion.BulkInster(cmd.CmdMapper.GetKeys())
	defer rl.Close()
	for {
		input, err := rl.Readline()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}
		trimedInput := strings.Split(strings.TrimSpace(input), " ")
		cmd, argv := trimedInput[0], trimedInput[1:]
		cmdExecuter.Execute(cmd, argv)
	}
}
