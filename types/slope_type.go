package types

type SlopeType int

const (
	StHorizontal SlopeType = iota
	StVertical             = iota
	StPositive             = iota
	StNegative             = iota
)
