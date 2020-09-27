package doom

import (
	"godoom/input"
	"godoom/video"
)

func bindVariables() {
	input.BindInputVariables()
	video.BindVideoVariables()

}
