package types

type Action struct {
	Void VoidAction
	// TODO ??
}

func NullAction() {}

type VoidAction func()
