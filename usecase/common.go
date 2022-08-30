package usecase

import (
	"github.com/bwmarrin/discordgo"
	"github.com/christiansoetanto/better-servant-of-servus-dei/config/configtypes"
	"log"
)

func (u *usecase) getGuildConfig(guildId string) (configtypes.GuildConfig, bool) {
	cfg, ok := u.Config.GuildConfig[guildId]
	return cfg, ok
}

func (u *usecase) isMod(userId string, guildCfg configtypes.GuildConfig) bool {
	_, ok := guildCfg.Moderator[userId]
	return ok
}

func (u *usecase) alertNonMod(i *discordgo.InteractionCreate) error {
	err := u.Session.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: NiceTryBro,
		},
	})
	if err != nil {
		return err
	}
	return nil
}

func (u *usecase) errorReporter(err error) {
	channel, e := u.Session.UserChannelCreate("255514888041005057")
	if e != nil {
		log.Print(e.Error())
		return
	}
	_, e = u.Session.ChannelMessageSend(channel.ID, err.Error())
	if e != nil {
		log.Print(e.Error())
		return
	}
	log.Print(err.Error())
}
