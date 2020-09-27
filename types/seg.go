package types

type Seg struct {
	V1 *Vertex
	V2 *Vertex

	Offset Fixed

	Angle Angle

	Sidedef *Side
	Linedef *Line

	// Sector references.
	// Could be retrieved from the linedef too.
	FrontSector *Sector
	// BackSector is null for one-sided lines
	BackSector *Sector
}
