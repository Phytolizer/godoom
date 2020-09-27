package main

import (
	"godoom/argv"
	"godoom/doom"
	"godoom/meta"
	"godoom/response_file"
	"os"
	sdl2 "src/github.com/veandco/go-sdl2/sdl"
)

func main() {
	argv.MyArgs = os.Args

	if argv.ParmExists("-version") || argv.ParmExists("--version") {
		println(meta.PackageString)
		os.Exit(0)
	}

	response_file.FindResponseFile()

	sdl2.SetHint(sdl2.HINT_NO_SIGNAL_HANDLERS, "1")

	doom.Main()
}
