package dbot

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/bwmarrin/discordgo"
	"github.com/christiansoetanto/better-servant-of-servus-dei/config"
	"log"
)

type Bot struct {
	Cfg             config.Config
	Session         *discordgo.Session
	FirestoreClient *firestore.Client
	DevMode         bool
}

const TOKEN = "OTc0MzExMDU5NjgwODIxMjY4.GukOAG.Cn99_DaXraufhv6m7CxoyXNgqQq7AmmqSIx0Qc"

func New(cfg config.Config, firestoreClient *firestore.Client, devMode bool) *Bot {
	return &Bot{
		Cfg:             cfg,
		FirestoreClient: firestoreClient,
		DevMode:         devMode,
	}
}

func (b *Bot) NewSession() (err error) {
	b.Session, err = discordgo.New("Bot " + TOKEN)
	return err
}

func (b *Bot) SetIntent() {
	b.Session.Identify.Intents = discordgo.IntentGuildMessages | discordgo.IntentGuildMessageReactions | discordgo.IntentDirectMessages
}

func (b *Bot) OpenConnection() error {
	return b.Session.Open()
}

func (b *Bot) LoadAllHandlers() {
	b.initReadyHandler()
	b.initMessageHandler()
	b.initReactionAddHandler()
}

func (b *Bot) InitAllCronJobs() {
	b.initCronJob()
}

func (b *Bot) TestFirestore(ctx context.Context) {
	ref, res, err := b.FirestoreClient.Collection("users").Add(ctx, map[string]interface{}{
		"first": "Ada",
		"last":  "Lovelace",
		"born":  1815,
	})
	_, _ = ref, res
	if err != nil {
		log.Fatalf("Failed adding alovelace: %v", err)
	}
}
