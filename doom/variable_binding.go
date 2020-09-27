package doom

import (
	"godoom/controls"
	"godoom/input"
	"godoom/joystick"
	"godoom/misc"
	"godoom/sound"
	"godoom/video"
)

func bindVariables() {
	input.BindInputVariables()
	video.BindVideoVariables()
	joystick.BindJoystickVariables()
	sound.BindSoundVariables()

	controls.BindBaseControls()
	controls.BindWeaponControls()
	controls.BindMapControls()
	controls.BindMenuControls()
	controls.BindChatControls(misc.MaxPlayers)
}
