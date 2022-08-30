package config

import (
	"github.com/christiansoetanto/better-servant-of-servus-dei/config/configtypes"
)

type Config struct {
	GuildConfig GuildConfigMap
	AppConfig   AppConfig
}
type AppConfig struct {
	DevMode            bool
	FirestoreProjectId string
}
type GuildConfigMap map[string]configtypes.GuildConfig

func Init(devMode bool) Config {
	config := Config{
		GuildConfig: InitGuildConfig(devMode),
		AppConfig: AppConfig{
			DevMode:            devMode,
			FirestoreProjectId: "youtube-title-updater-340409",
		},
	}
	return config
}

func InitGuildConfig(devMode bool) GuildConfigMap {
	if devMode {
		return initDevGuildConfig()
	}
	servusDeiConfig := getServusDeiGuildConfig()
	devConfig := getDevGuildConfig()
	cfg := GuildConfigMap{
		servusDeiConfig.Guild.GuildID: servusDeiConfig,
		devConfig.Guild.GuildID:       devConfig,
	}
	return cfg

}
func initDevGuildConfig() GuildConfigMap {
	devConfig := getDevGuildConfig()
	cfg := GuildConfigMap{
		devConfig.Guild.GuildID: devConfig,
	}
	return cfg

}
