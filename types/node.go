package types

// BSP node
type Node struct {
	// Partition line
	X  Fixed
	Y  Fixed
	Dx Fixed
	Dy Fixed

	// Bounding box for each child
	BBox [2][4]Fixed

	// If NF_SUBSECTOR, it's a subsector.
	Children [2]uint16
}
