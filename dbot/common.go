package dbot

import (
	"fmt"
	"github.com/christiansoetanto/better-servant-of-servus-dei/config/configtypes"
	"log"
)

func (b *Bot) getGuildConfig(guildId string) (configtypes.GuildConfig, error) {
	cfg, ok := b.Cfg[guildId]
	if !ok {
		return configtypes.GuildConfig{}, fmt.Errorf("guild %s not found", guildId)
	}
	return cfg, nil
}

func (b *Bot) isMod(userId string, guildCfg configtypes.GuildConfig) bool {
	_, ok := guildCfg.Moderator[userId]
	return ok
}

func (b *Bot) errorReporter(err error) {
	channel, e := b.Session.UserChannelCreate("255514888041005057")
	if e != nil {
		log.Print(e.Error())
		return
	}
	_, e = b.Session.ChannelMessageSend(channel.ID, err.Error())
	if e != nil {
		log.Print(e.Error())
		return
	}
	log.Print(err.Error())
}
