package config

import (
	"github.com/christiansoetanto/better-servant-of-servus-dei/config/configtypes"
)

type Config map[string]configtypes.GuildConfig

func Init(devMode bool) Config {
	if devMode {
		return initDev()
	}
	servusDeiConfig := getServusDeiConfig()
	devConfig := getDevConfig()
	cfg := Config{
		servusDeiConfig.Guild.GuildID: servusDeiConfig,
		devConfig.Guild.GuildID:       devConfig,
	}
	return cfg

}
func initDev() Config {
	devConfig := getDevConfig()
	cfg := Config{
		devConfig.Guild.GuildID: devConfig,
	}
	return cfg

}
