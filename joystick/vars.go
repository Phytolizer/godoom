package joystick

const (
	NumVirtualButtons = 11
)

var (
	UseJoystick  = false
	Guid         = ""
	Index        = -1
	XAxis        = 0
	YAxis        = 1
	StrafeAxis   = -1
	XInvert      = false
	YInvert      = false
	StrafeInvert = false
	LookAxis     = -1
	LookInvert   = false

	PhysicalButtons = [NumVirtualButtons]int{
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	}
)
