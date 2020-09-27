package configuration

import (
	"godoom/errors"
	"unsafe"
)

func interfacify(x *interface{}) *interface{} {
	return x
}

func BindIntVariable(key string, value *int) {
	variable := GetDefaultForName(key)
	if variable.Kind != DefaultInt && variable.Kind != DefaultIntHex && variable.Kind != DefaultKey {
		errors.Error("Invalid type for variable ", key, " (attempted to bind it to integer)")
	}
	variable.Location = unsafe.Pointer(value)
	variable.Bound = true
}

func BindBoolVariable(key string, value *bool) {
	variable := GetDefaultForName(key)
	if variable.Kind != DefaultBool {
		errors.Error("Invalid type for variable ", key, " (attempted to bind it to boolean)")
	}
	variable.Location = unsafe.Pointer(value)
	variable.Bound = true
}

func BindStringVariable(key string, value *string) {
	variable := GetDefaultForName(key)
	if variable.Kind != DefaultString {
		errors.Error("Invalid type for variable ", key, " (attempted to bind it to string)")
	}
	variable.Location = unsafe.Pointer(value)
	variable.Bound = true
}

func BindFloatVariable(key string, value *float64) {
	variable := GetDefaultForName(key)
	if variable.Kind != DefaultFloat {
		errors.Error("Invalid type for variable ", key, " (attempted to bind it to float)")
	}
	variable.Location = unsafe.Pointer(value)
	variable.Bound = true
}
