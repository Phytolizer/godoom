package types

// A linedef
type Line struct {
	// Vertices, from V1 to V2
	V1 *Vertex
	V2 *Vertex

	// Precalculated V2 - V1 for side checking
	Dx Fixed
	Dy Fixed

	// Animation related
	Flags   int16
	Special int16
	Tag     int16

	// Visual appearance: Sidedefs. (Side)
	// SideNum[1] will be -1 if one-sided
	SideNum [2]int16

	// Another bounding box, for the extent of this linedef
	BBox [4]Fixed

	// To aid move clipping
	SlopeType SlopeType

	// front and back sector
	//
	// note: redundant? can be retrieved from sidedefs...
	FrontSector *Sector
	BackSector  *Sector

	// if == validcount, already checked :)
	ValidCount int

	// Thinker for reversible actions
	SpecialData *uint8
}
