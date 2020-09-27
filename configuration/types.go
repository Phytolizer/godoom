package configuration

import "unsafe"

type DefaultKind int

const (
	DefaultInt    DefaultKind = iota
	DefaultIntHex             = iota
	DefaultString             = iota
	DefaultFloat              = iota
	DefaultKey                = iota
	DefaultBool               = iota
)

type Default struct {
	Name string
	// ptr to bool, int, string, or float
	Location unsafe.Pointer
	// tells what type Location points to
	Kind DefaultKind
	// if it's a key value, the original scancode from the cfg file is here
	//
	// if 0, it wasn't from a cfg file
	Untranslated int
	// translated value from config file on startup, see Untranslated
	OriginalTranslated int
	// bound to a variable and being used?
	Bound bool
}

type DefaultCollection struct {
	Defaults []Default
	FileName string
}
