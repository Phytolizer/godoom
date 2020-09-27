package doom

import (
	"godoom/configuration"
	"godoom/controls"
	"godoom/doomstat"
	"godoom/english"
	"godoom/game"
	"godoom/hud"
	"godoom/input"
	"godoom/joystick"
	"godoom/menu"
	"godoom/misc"
	"godoom/net"
	"godoom/sound"
	"godoom/video"
	"strconv"
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

	net.BindVariables()

	configuration.BindIntVariable("mouse_sensitivity", &doomstat.MouseSensitivity)
	configuration.BindIntVariable("sfx_volume", &doomstat.SfxVolume)
	configuration.BindIntVariable("music_volume", &doomstat.MusicVolume)
	configuration.BindBoolVariable("show_messages", &menu.ShowMessages)
	configuration.BindIntVariable("screenblocks", &menu.ScreenBlocks)
	configuration.BindIntVariable("detaillevel", &menu.DetailLevel)
	configuration.BindIntVariable("snd_channels", &sound.Channels)
	configuration.BindBoolVariable("vanilla_savegame_limit", &game.VanillaSavegameLimit)
	configuration.BindBoolVariable("vanilla_demo_limit", &game.VanillaDemoLimit)
	configuration.BindBoolVariable("show_endoom", &showEndoom)
	configuration.BindBoolVariable("show_diskicon", &showDiskIcon)

	for i := 0; i < 10; i += 1 {
		hud.ChatMacros[i] = hud.ChatMacroDefaults[i]
		buf := "chatmacro"
		buf += strconv.FormatInt(int64(i), 10)
		configuration.BindStringVariable(buf, &hud.ChatMacros[i])
	}
}
