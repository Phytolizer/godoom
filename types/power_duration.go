package types

import "godoom/timer"

type PowerDuration int

const (
	InvulnTics = PowerDuration(30 * timer.TicRate)
	InvisTics  = PowerDuration(60 * timer.TicRate)
	InfraTics  = PowerDuration(120 * timer.TicRate)
	IronTics   = PowerDuration(60 * timer.TicRate)
)
