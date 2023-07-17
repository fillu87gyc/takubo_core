package lib

type color int

const (
	Reset color = iota
	Black
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
)
const (
	reset   = "\033[0m"
	black   = "\033[30m"
	red     = "\033[31m"
	green   = "\033[32m"
	yellow  = "\033[33m"
	blue    = "\033[34m"
	magenta = "\033[35m"
	cyan    = "\033[36m"
	white   = "\033[37m"
)

func Color(text string, colorCode color) string {
	c := ""
	switch colorCode {
	case Reset:
		c = reset
	case Black:
		c = black
	case Red:
		c = red
	case Green:
		c = green
	case Yellow:
		c = yellow
	case Blue:
		c = blue
	case Magenta:
		c = magenta
	case Cyan:
		c = cyan
	case White:
		c = white
	}
	return c + text + reset
}
