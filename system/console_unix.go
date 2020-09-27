// +build !windows

package system

import (
	"os"
	"src/github.com/mattn/go-isatty"
)

func ConsoleStdout() bool {
	return isatty.IsTerminal(os.Stdout.Fd())
}
