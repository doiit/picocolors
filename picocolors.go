package picocolors

import (
	"fmt"
	"os"
	"strings"

	"github.com/mattn/go-isatty"
)

var SUPPORT_COLOR = isColorSupported()

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func isColorSupported() bool {
	noColor := os.Getenv("NO_COLOR") != "" || contains(os.Args, "--no-color")
	forceColor := os.Getenv("FORCE_COLOR") != "" || os.Getenv("CLICOLOR_FORCE") != "" || contains(os.Args, "--color")
	terminal := isatty.IsTerminal(os.Stdout.Fd()) || isatty.IsCygwinTerminal(os.Stdout.Fd()) || os.Getenv("TERM") == "dumb"
	ci := os.Getenv("CI") != ""

	return !noColor && (forceColor || terminal || ci)
}

type callFunc = func(input interface{}) string

func replaceClose(str, close, replace string, index int) string {
	start := str[0:index]
	end := str[index+len(close):]
	nextIndex := strings.Index(end, close)
	if nextIndex == -1 {
		return start + end
	}

	return replaceClose(end, close, replace, nextIndex)
}

func formater(open, close, replace string) callFunc {
	return func(input interface{}) string {
		str := input.(string)
		index := strings.Index(str[len(open):], close)
		if index != -1 {
			index += len(open)
			return open + replaceClose(str, close, replace, index) + close
		}

		return open + str + close
	}
}

func Reset(input interface{}) string {
	if SUPPORT_COLOR {
		return fmt.Sprintf("\x1b[0m%v\x1b[0m", input)
	}
	return input.(string)
}
func Bold(input interface{}) string {
	if SUPPORT_COLOR {
		return formater("\x1b[1m", "\x1b[22m", "\x1b[22m\x1b[1m")(input)
	}
	return input.(string)
}
func Dim(input interface{}) string {
	if SUPPORT_COLOR {
		return formater("\x1b[2m", "\x1b[22m", "\x1b[22m\x1b[2m")(input)
	}
	return input.(string)
}
func Italic(input interface{}) string {
	if SUPPORT_COLOR {
		return formater("\x1b[3m", "\x1b[23m", "")(input)
	}
	return input.(string)
}
func Underline(input interface{}) string {
	if SUPPORT_COLOR {
		return formater("\x1b[4m", "\x1b[24m", "")(input)
	}
	return input.(string)
}
func Inverse(input interface{}) string {
	if SUPPORT_COLOR {
		return formater("\x1b[7m", "\x1b[27m", "")(input)
	}
	return input.(string)
}
func Hidden(input interface{}) string {
	if SUPPORT_COLOR {
		return formater("\x1b[8m", "\x1b[28m", "")(input)
	}
	return input.(string)
}
func Strikethrough(input interface{}) string {
	if SUPPORT_COLOR {
		return formater("\x1b[9m", "\x1b[29m", "")(input)
	}
	return input.(string)
}
func Black(input interface{}) string {
	if SUPPORT_COLOR {
		return formater("\x1b[30m", "\x1b[39m", "")(input)
	}
	return input.(string)
}
func Red(input interface{}) string {
	if SUPPORT_COLOR {
		return formater("\x1b[31m", "\x1b[39m", "")(input)
	}
	return input.(string)
}
func Green(input interface{}) string {
	if SUPPORT_COLOR {
		return formater("\x1b[32m", "\x1b[39m", "")(input)
	}
	return input.(string)
}
func Yellow(input interface{}) string {
	if SUPPORT_COLOR {
		return formater("\x1b[33m", "\x1b[39m", "")(input)
	}
	return input.(string)
}
func Blue(input interface{}) string {
	if SUPPORT_COLOR {
		return formater("\x1b[34m", "\x1b[39m", "")(input)
	}
	return input.(string)
}
func Magenta(input interface{}) string {
	if SUPPORT_COLOR {
		return formater("\x1b[35m", "\x1b[39m", "")(input)
	}
	return input.(string)
}
func Cyan(input interface{}) string {
	if SUPPORT_COLOR {
		return formater("\x1b[36m", "\x1b[39m", "")(input)
	}
	return input.(string)
}
func White(input interface{}) string {
	if SUPPORT_COLOR {
		return formater("\x1b[37m", "\x1b[39m", "")(input)
	}
	return input.(string)
}
func Gray(input interface{}) string {
	if SUPPORT_COLOR {
		return formater("\x1b[90m", "\x1b[39m", "")(input)
	}
	return input.(string)
}
func BgBlack(input interface{}) string {
	if SUPPORT_COLOR {
		return formater("\x1b[40m", "\x1b[49m", "")(input)
	}
	return input.(string)
}
func BgRed(input interface{}) string {
	if SUPPORT_COLOR {
		return formater("\x1b[41m", "\x1b[49m", "")(input)
	}
	return input.(string)
}
func BgGreen(input interface{}) string {
	if SUPPORT_COLOR {
		return formater("\x1b[42m", "\x1b[49m", "")(input)
	}
	return input.(string)
}
func BgYellow(input interface{}) string {
	if SUPPORT_COLOR {
		return formater("\x1b[43m", "\x1b[49m", "")(input)
	}
	return input.(string)
}
func BgBlue(input interface{}) string {
	if SUPPORT_COLOR {
		return formater("\x1b[44m", "\x1b[49m", "")(input)
	}
	return input.(string)
}
func BgMagenta(input interface{}) string {
	if SUPPORT_COLOR {
		return formater("\x1b[45m", "\x1b[49m", "")(input)
	}
	return input.(string)
}
func BgCyan(input interface{}) string {
	if SUPPORT_COLOR {
		return formater("\x1b[46m", "\x1b[49m", "")(input)
	}
	return input.(string)
}
func BgWhite(input interface{}) string {
	if SUPPORT_COLOR {
		return formater("\x1b[47m", "\x1b[49m", "")(input)
	}
	return input.(string)
}
