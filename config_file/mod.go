package config_file

import (
	"godoom/meta"
	"godoom/misc"

	sdl2 "github.com/veandco/go-sdl2/sdl"
)

var configDir = ""
var defaultMainConfig = ""
var defaultExtraConfig = ""

func getDefaultConfigDir() string {
	return sdl2.GetPrefPath("", meta.PackageTarname)
}

func SetConfigDir(dir []byte) {
	if dir != nil {
		configDir = string(dir)
	} else {
		configDir = getDefaultConfigDir()
	}

	if len(configDir) > 0 {
		println("Using ", configDir, " for configuration and saves.")
	}

	misc.MakeDirectory(configDir)
}

func SetConfigFilenames(mainConfig string, extraConfig string) {
	defaultMainConfig = mainConfig
	defaultExtraConfig = extraConfig
}
