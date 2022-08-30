package dbot

import (
	"github.com/bwmarrin/discordgo"
	"log"
)

const TOKEN = "OTc0MzExMDU5NjgwODIxMjY4.GukOAG.Cn99_DaXraufhv6m7CxoyXNgqQq7AmmqSIx0Qc"

func New() *discordgo.Session {
	session, err := discordgo.New("Bot " + TOKEN)
	if err != nil {
		log.Fatal(err)
	}
	return session
}
