package types

import "math"

type Fixed int

const (
	FracBits = 16
	FracUnit = 1 << FracBits
)

func FixedMul(a Fixed, b Fixed) Fixed {
	return Fixed((int64(a) * int64(b)) >> FracBits)
}

func abs(x Fixed) Fixed {
	if x < 0 {
		return -x
	}
	return x
}

func FixedDiv(a Fixed, b Fixed) Fixed {
	if (abs(a) >> 14) >= Fixed(abs(b)) {
		if a^b < 0 {
			return math.MinInt32
		}
		return math.MaxInt32
	}
	return Fixed((int64(a) << FracBits) / int64(b))
}
