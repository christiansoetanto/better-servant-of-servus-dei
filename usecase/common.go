package usecase

import (
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/christiansoetanto/better-servant-of-servus-dei/config/configtypes"
	"log"
	"runtime"
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

type ErrorReporter struct {
	ErrorDetail string `json:"error_detail"`
	File        string `json:"file"`
	LineNumber  int    `json:"line_number"`
	FuncName    string `json:"func_name"`
}

func jsonEscape(i string) string {
	b, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}
	s := string(b)
	return s[1 : len(s)-1]
}

func (u *usecase) errorReporter(err error) {
	channel, e := u.Session.UserChannelCreate("255514888041005057")
	if e != nil {
		log.Print(e.Error())
		return
	}
	errorReporter := ErrorReporter{
		ErrorDetail: jsonEscape(err.Error()),
	}
	pc, file, line, ok := runtime.Caller(1)
	if ok {
		details := runtime.FuncForPC(pc)
		if details != nil {
			errorReporter.FuncName = details.Name()
			file, line := details.FileLine(pc)
			errorReporter.File = file
			errorReporter.LineNumber = line
		} else {
			errorReporter.File = file
			errorReporter.LineNumber = line
		}
	}

	rep, _ := json.Marshal(errorReporter)

	_, e = u.Session.ChannelMessageSend(channel.ID, fmt.Sprintf("%s\n", string(rep)))
	if e != nil {
		log.Print(e.Error())
		return
	}
	log.Print(err.Error())
}
