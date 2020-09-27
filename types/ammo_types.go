package types

type AmmoType int

const (
	AmClip  AmmoType = iota
	AmShell          = iota
	AmCell           = iota
	AmMisl           = iota
	NumAmmo          = iota
	NoAmmo           = iota
)
