package usecase

import (
	"github.com/bwmarrin/discordgo"
	"log"
)

func (u *usecase) initReadyHandler() {
	u.Session.AddHandler(u.readyHandler)
}

func (u *usecase) readyHandler(s *discordgo.Session, r *discordgo.Ready) {
	log.Printf("Logged in as: %v#%v id: %v", s.State.User.Username, s.State.User.Discriminator, s.State.User.ID)
}
