package readline

import (
	"fmt"
	"os"

	"github.com/pkg/term"
)

// https://pkg.go.dev/golang.org/x/term#pkg-variables alt to this term ?
// inspi https://github.com/ollama/ollama/blob/main/readline/buffer.go
type termType string
type termMode string

type Instance struct {
	term       *term.Term
	buffer     []rune
	Completion *Completion
}

func New(tType termType, tMode termMode) *Instance {
	t, err := term.Open(string(tType))
	if err != nil {
		panic("Unable to open the terminal")
	}
	if tMode == RAW {
		err := term.RawMode(t)
		if err != nil {
			panic("Unable to put the terminal in raw mode")
		}
	} else if tMode == CBREAK {
		err := term.CBreakMode(t)
		if err != nil {
			panic("Unable to put the terminal in cbreak mode")
		}
	}
	return &Instance{
		term:       t,
		Completion: NewCompletion(true),
	}
}
func (i *Instance) Readline() (string, error) {
	fmt.Fprint(os.Stdout, "$ ")
	for {
		var read int
		var err error
		readBytes := make([]byte, 3)
		read, err = i.term.Read(readBytes)
		if err != nil {
			return "", err
		}
		if read == 3 {
			// Arrow keys are prefixed with the ANSI escape code which take up the first two bytes.
			// The third byte is the key specific value we are looking for.
			// For example the left arrow key is '<esc>[A' while the right is '<esc>[C'
			// See: https://en.wikipedia.org/wiki/ANSI_escape_code
			// A ctrl key (up,down,lef,right) is 3 bytes length and the specific key is in the third byte.
			//TODO handle arrow key
			continue
		}
		r := readBytes[0]
		//https://en.wikipedia.org/wiki/List_of_Unicode_characters
		switch r {
		case enter:
			fmt.Print("\n")
			return string(i.buffer), nil
		case delete:
			if len(i.buffer) > 0 {
				i.buffer = i.buffer[:len(i.buffer)-1]
				fmt.Fprint(os.Stdout, ("\b \b")) // Move cursor back, print space, move cursor back again
			}
		case tab:
			if len(i.buffer) == 0 {
				fmt.Print("\t")
			} else {
				prefix := string(i.buffer)
				suggestions := i.Completion.GetSuggestions(prefix)
				if len(suggestions) > 0 {
					i.buffer = []rune(suggestions[0])
					fmt.Print("\r$ " + string(i.buffer))
				}
			}

		default:
			if r >= space && r < delete {
				i.buffer = append(i.buffer, rune(r))
				fmt.Fprint(os.Stdout, string(r))
			}
		}
	}

}

func (i *Instance) Close() error {
	i.term.Restore()
	i.Close()
	return nil
}
func (i *Instance) emptyBuffer() {

}
func (i *Instance) ClearScrean() {
	// \033[H : Move the cursor to the top left corner
	// \033[2J : Clear the screen
	fmt.Fprintln(os.Stdout, "\033[H\033[2J")
}
func (i *Instance) ShowCursor() {
	fmt.Fprintln(os.Stdout, "\033[?25h")
}
