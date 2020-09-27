package types

type WeaponType int

const (
	WpFist         WeaponType = iota
	WpPistol                  = iota
	WpShotgun                 = iota
	WpChaingun                = iota
	WpMissile                 = iota
	WpPlasma                  = iota
	WpBfg                     = iota
	WpChainsaw                = iota
	WpSuperShotgun            = iota
	NumWeapons                = iota
	WpNoChange                = iota
)
