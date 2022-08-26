package dbot

import (
	"github.com/bwmarrin/discordgo"
	"log"
)

func (b *Bot) readyHandler() func(s *discordgo.Session, r *discordgo.Ready) {
	return func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
	}
}
