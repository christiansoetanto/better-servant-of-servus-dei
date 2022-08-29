package dbot

import (
	"github.com/bwmarrin/discordgo"
	"log"
)

func (b *Bot) initReadyHandler() {
	b.Session.AddHandler(b.readyHandler1)
}

func (b *Bot) readyHandler1(s *discordgo.Session, r *discordgo.Ready) {
	log.Printf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
}
