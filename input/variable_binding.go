package input

import "godoom/configuration"

func BindInputVariables() {
	configuration.BindFloatVariable("mouse_acceleration", &MouseAcceleration)
	configuration.BindIntVariable("mouse_threshold", &MouseThreshold)
	configuration.BindBoolVariable("vanilla_keyboard_mapping", &VanillaKeyboardMapping)
	configuration.BindBoolVariable("novert", &NoVert)
}
