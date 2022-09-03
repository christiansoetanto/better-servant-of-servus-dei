package dbot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
)

func New() *discordgo.Session {
	session, err := discordgo.New(fmt.Sprintf("Bot %s", os.Getenv("BOTTOKEN")))
	if err != nil {
		log.Fatal(err)
	}
	return session
}
