package doom

import (
	"godoom/argv"
	"godoom/deh"
	"godoom/video"
	"godoom/wad"
)

var showEndoom = 1
var mainLoopStarted = false

func endoom() {
	if showEndoom == 0 || !mainLoopStarted || video.ScreensaverMode || argv.ParmExists("-testcontrols") {
		return
	}

	endoom := wad.CacheLumpName(deh.String("ENDOOM"))
	endoomInternal(endoom)
}

func endoomInternal(endoomData []byte) {
	// TODO
}
