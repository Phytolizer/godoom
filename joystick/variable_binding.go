package joystick

import (
	"godoom/configuration"
	"strconv"
)

func BindJoystickVariables() {
	configuration.BindBoolVariable("use_joystick", &UseJoystick)
	configuration.BindStringVariable("joystick_guid", &Guid)
	configuration.BindIntVariable("joystick_index", &Index)
	configuration.BindIntVariable("joystick_x_axis", &XAxis)
	configuration.BindIntVariable("joystick_y_axis", &YAxis)
	configuration.BindIntVariable("joystick_strafe_axis", &StrafeAxis)
	configuration.BindBoolVariable("joystick_x_invert", &XInvert)
	configuration.BindBoolVariable("joystick_y_invert", &YInvert)
	configuration.BindBoolVariable("joystick_strafe_invert", &StrafeInvert)
	configuration.BindIntVariable("joystick_look_axis", &LookAxis)
	configuration.BindBoolVariable("joystick_look_invert", &LookInvert)

	for i := 0; i < NumVirtualButtons; i += 1 {
		name := "joystick_physical_button" + strconv.FormatInt(int64(i), 10)
		configuration.BindIntVariable(name, &PhysicalButtons[i])
	}
}
