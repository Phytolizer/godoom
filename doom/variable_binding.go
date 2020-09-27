package doom

import (
	"godoom/controls"
	"godoom/english"
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

	controls.KeyMultiMsgPlayer[0] = english.HuStrKeyGreen
	controls.KeyMultiMsgPlayer[1] = english.HuStrKeyIndigo
	controls.KeyMultiMsgPlayer[2] = english.HuStrKeyBrown
	controls.KeyMultiMsgPlayer[3] = english.HuStrKeyRed
}
