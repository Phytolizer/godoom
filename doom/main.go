package doom

import (
	"godoom/argv"
	"godoom/config_file"
	"godoom/deh"
	"godoom/doomstat"
	"godoom/english"
	"godoom/game"
	"godoom/meta"
	"godoom/net"
	"godoom/system"
	"godoom/types"
	"godoom/video"
	"os"
	"strconv"
)

func Main() {
	system.AtExit(endoom, false)
	system.PrintBanner(meta.PackageString)

	if argv.ParmExists("-dedicated") {
		println("Dedicated server mode.")
		net.DedicatedServer()

		// this is unreachable
		os.Exit(-1)
	}

	if argv.ParmExists("-search") {
		net.MasterQuery()
		os.Exit(0)
	}

	p := argv.CheckParmWithArgs("-query", 1)
	if p != 0 {
		net.QueryAddress(argv.MyArgs[p+1])
	}

	if argv.ParmExists("-localsearch") {
		net.LanQuery()
	}

	doomstat.NoMonsters = argv.ParmExists("-nomonsters")
	doomstat.RespawnParm = argv.ParmExists("-respawn")
	doomstat.FastParm = argv.ParmExists("-fast")
	doomstat.DevParm = argv.ParmExists("-devparm")
	video.DisplayFpsDots(doomstat.DevParm)
	if argv.ParmExists("-deathmatch") {
		doomstat.Deathmatch = 1
	}
	if argv.ParmExists("-altdeath") {
		doomstat.Deathmatch = 2
	}

	if doomstat.DevParm {
		deh.Print(english.DevStr)
	}

	config_file.SetConfigDir(nil)

	p = argv.CheckParm("-turbo")
	if p != 0 {
		scale := types.Fixed(200)
		if p < len(argv.MyArgs)-1 {
			scale, err := strconv.Atoi(argv.MyArgs[p+1])
			if err != nil {
				scale = 10
			}
			if scale < 10 {
				scale = 10
			}
			if scale > 400 {
				scale = 400
			}
		}
		deh.Print("turbo scale: %i%%\n", scale)
		// not using *= to avoid precision loss
		game.ForwardMove[0] = game.ForwardMove[0] * scale / 100
		game.ForwardMove[1] = game.ForwardMove[1] * scale / 100
		game.SideMove[0] = game.SideMove[0] * scale / 100
		game.SideMove[1] = game.SideMove[1] * scale / 100
	}

	config_file.SetConfigFilenames("default.cfg", meta.PackageTarname+".cfg")
	bindVariables()
}
