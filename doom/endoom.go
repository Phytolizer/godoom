package doom

import (
	"godoom/argv"
	"godoom/deh"
	"godoom/video"
	"godoom/wad"
)

func endoom() {
	if !showEndoom || !mainLoopStarted || video.ScreensaverMode || argv.ParmExists("-testcontrols") {
		return
	}

	endoom := wad.CacheLumpName(deh.String("ENDOOM"))
	endoomInternal(endoom)
}

func endoomInternal(endoomData []byte) {
	// TODO
}
