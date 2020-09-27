package hud

import "godoom/english"

// Variables that can actually change
var (
	ChatMacros [10]string
)

// "Constants", no const arrays in Go :(
var (
	ChatMacroDefaults = [...]string{
		english.HuStrChatMacro0,
		english.HuStrChatMacro1,
		english.HuStrChatMacro2,
		english.HuStrChatMacro3,
		english.HuStrChatMacro4,
		english.HuStrChatMacro5,
		english.HuStrChatMacro6,
		english.HuStrChatMacro7,
		english.HuStrChatMacro8,
		english.HuStrChatMacro9,
	}
)
