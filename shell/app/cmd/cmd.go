package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/codecrafters-io/shell-starter-go/app/utils"
)

type cmdMapper map[string]func(argv []string, stderr *os.File, stdout *os.File)

func (cmdMapper) GetKeys() []string {
	keys := make([]string, 0, len(CmdMapper))
	for k := range CmdMapper {
		keys = append(keys, k)
	}
	return keys
}

var CmdMapper cmdMapper

func Type(argv []string, stderr *os.File, stdout *os.File) {
	if len(argv) == 0 {
		return
	}
	//Priority check for builtin commands
	if _, found := CmdMapper[argv[0]]; found {
		fmt.Fprintln(stdout, argv[0]+" is a shell builtin")
		return
	}
	path, found := utils.FindPath(argv[0])
	if found {
		fmt.Fprintln(stdout, argv[0]+" is "+path)
		return
	}
	fmt.Fprintln(stdout, argv[0]+": not found")
}

type CmdExecuter struct {
	stderr *os.File
	stdout *os.File
	stdin  *os.File
	mapper cmdMapper
}

func New(stderr *os.File, stdout *os.File, stdin *os.File) *CmdExecuter {
	return &CmdExecuter{
		stderr: stderr,
		stdout: stdout,
		stdin:  stdin,
		mapper: CmdMapper,
	}
}

func (c *CmdExecuter) Execute(cmd string, argv []string) {
	execBuiltInCmd, builtInFound := c.mapper[cmd]
	if !builtInFound {
		_, found := utils.FindPath(cmd)
		if found {
			cmd := exec.Command(cmd, argv...)
			cmd.Stdout = c.stdout
			cmd.Stderr = c.stderr
			cmd.Stdin = c.stdin
			err := cmd.Run()
			if err != nil {
				fmt.Fprintln(c.stderr, "Error executing command: ", err)
			}
		} else {
			fmt.Fprintln(c.stdout, cmd+": command not found")
		}
	} else {
		execBuiltInCmd(argv, c.stderr, c.stdout)
	}
}

// Need to init it at the package initilization to avoid init cycle as the type function is using it
func init() {
	CmdMapper = map[string]func(argv []string, stderr *os.File, stdout *os.File){
		"echo": Echo,
		"exit": Exit,
		"type": Type,
		"pwd":  Pwd,
		"cd":   ChangeDirectory,
	}
}
