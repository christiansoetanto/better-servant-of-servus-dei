package dbot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/christiansoetanto/better-servant-of-servus-dei/config"
)

type Bot struct {
	Cfg     config.Config
	Session *discordgo.Session
}

const TOKEN = "OTc0MzExMDU5NjgwODIxMjY4.GgLLqU.Sq8jkCbzRYS-4uQVa3U31V9pgvQNjCrKWvKLtA"

func New(cfg config.Config) *Bot {
	return &Bot{
		Cfg: cfg,
	}
}

func (b *Bot) NewSession() (err error) {
	b.Session, err = discordgo.New("Bot " + TOKEN)
	return err
}

func (b *Bot) OpenConnection() error {
	return b.Session.Open()
}

func (b *Bot) LoadAllHandlers() {
	b.Session.AddHandler(b.readyHandler())
	b.Session.AddHandler(b.messageHandler())
}
