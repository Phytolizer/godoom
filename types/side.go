package types

type Side struct {
	TextureOffset Fixed

	RowOffset Fixed

	TopTexture    int16
	BottomTexture int16
	MidTexture    int16

	Sector *Sector
}
