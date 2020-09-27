package doom

import (
	"godoom/input"
	"godoom/joystick"
	"godoom/sound"
	"godoom/video"
)

func bindVariables() {
	input.BindInputVariables()
	video.BindVideoVariables()
	joystick.BindJoystickVariables()
	sound.BindSoundVariables()
}
