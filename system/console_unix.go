// +build !windows

package system

import (
	"os"

	"github.com/mattn/go-isatty"
)

func ConsoleStdout() bool {
	return isatty.IsTerminal(os.Stdout.Fd())
}
