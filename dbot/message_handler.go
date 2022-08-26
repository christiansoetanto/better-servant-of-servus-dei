package dbot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func (b *Bot) messageHandler() func(s *discordgo.Session, m *discordgo.MessageCreate) {
	return func(s *discordgo.Session, m *discordgo.MessageCreate) {
		fmt.Println(m.Content)
	}
}
