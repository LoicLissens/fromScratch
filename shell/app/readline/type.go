package readline

const (
	TTY     termType = "/dev/tty"
	RAW     termMode = "RAW"
	COOCKED termMode = "COOCKED"
	CBREAK  termMode = "CBREAK"
)

var (
	up             byte = 65
	down           byte = 66
	escape         byte = 27
	carriageReturn byte = 13
	tab            byte = 9
	space          byte = 32
	delete         byte = 127
	enter          byte = 10
)
var keys = map[byte]bool{
	up:   true,
	down: true,
}
