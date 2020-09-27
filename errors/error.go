package errors

import (
	"fmt"
	"godoom/argv"
	"godoom/meta"
	"godoom/system"
	"os"
	sdl2 "src/github.com/veandco/go-sdl2/sdl"
)

var alreadyQuitting = false

func Error(args ...interface{}) {
	if alreadyQuitting {
		_, _ = fmt.Fprint(os.Stderr, "Warning: recursive call to godoom.errors.Error detected...\n")
		os.Exit(-1)
	}
	alreadyQuitting = true
	// discarding errors, if this fails I really don't care
	_, _ = fmt.Fprint(os.Stderr, args...)
	_, _ = fmt.Fprint(os.Stderr, "\n\n")

	msgbuf := fmt.Sprint(args...)

	for entry := system.ExitFuncs; entry != nil; entry = entry.Next {
		if entry.RunOnError {
			entry.Func()
		}
	}

	exitGuiPopup := !argv.ParmExists("-nogui")
	if exitGuiPopup && !system.ConsoleStdout() {
		// ignore errors showing it, we tried ok?
		_ = sdl2.ShowSimpleMessageBox(sdl2.MESSAGEBOX_ERROR, meta.PackageString, msgbuf, nil)
	}

	sdl2.Quit()
	os.Exit(-1)
}
