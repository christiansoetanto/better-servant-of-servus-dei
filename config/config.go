package config

import (
	"github.com/christiansoetanto/better-servant-of-servus-dei/config/configtypes"
)

type Config map[configtypes.GuildID]configtypes.GuildConfig

func Init() Config {
	servusDeiConfig := getServusDeiConfig()
	devConfig := getDevConfig()
	cfg := Config{
		servusDeiConfig.Guild.GuildID: servusDeiConfig,
		devConfig.Guild.GuildID:       devConfig,
	}
	return cfg

}
