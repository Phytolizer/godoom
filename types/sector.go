package types

type Sector struct {
	FloorHeight   Fixed
	CeilingHeight Fixed
	FloorPic      int16
	CeilingPic    int16
	LightLevel    int16
	Special       int16
	Tag           int16

	SoundTraversed int

	SoundTarget *Mobj

	BlockBox [4]int

	SoundOrg DegenMobj

	ValidCount int

	ThingList *Mobj

	SpecialData *uint8

	LineCount int
	Lines     **Line
}
