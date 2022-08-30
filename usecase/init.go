package usecase

import (
	"context"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/christiansoetanto/better-servant-of-servus-dei/config"
	"github.com/christiansoetanto/better-servant-of-servus-dei/provider"
	"github.com/robfig/cron/v3"
	"sync"
)

type Resource struct {
	Provider provider.Provider
	Config   config.Config
	Session  *discordgo.Session
}

type usecase struct {
	*Resource
}

type Usecase interface {
	Init(ctx context.Context) error
	OpenSessionConnection() error
	CloseSessionConnection() error
	DoRemoveSlashCommand() error
	DoHelloWorld(ctx context.Context)
}

var obj Usecase
var once sync.Once

func GetUsecaseObject(resource *Resource) Usecase {
	once.Do(func() {
		obj = &usecase{
			Resource: resource,
		}
	})
	return obj
}

const DailyCron = "@daily"
const Every5SecondCron = "@every 5s"

func (u *usecase) Init(ctx context.Context) error {

	//handlers => open conn => cron jobs
	u.LoadAllHandlers()

	err := u.OpenSessionConnection()
	if err != nil {
		return err
	}
	err = u.DoRemoveSlashCommand()
	if err != nil {
		return err
	}
	u.registerSlashCommand()
	u.LoadAllCronJobs()

	return nil
}
func (u *usecase) OpenSessionConnection() error {
	u.Session.Identify.Intents = discordgo.IntentGuildMessages | discordgo.IntentGuildMessageReactions | discordgo.IntentDirectMessages
	return u.Session.Open()
}
func (u *usecase) CloseSessionConnection() error {
	fmt.Printf("session closed")
	return u.Session.Close()
}

func (u *usecase) LoadAllHandlers() {
	u.initReadyHandler()
	u.initMessageHandler()
	u.initReactionAddHandler()
	u.initCommandHandler()
}

func (u *usecase) LoadAllCronJobs() {
	c := cron.New()
	_, err := c.AddFunc(DailyCron, u.liturgicalCalendarCronJob)
	if err != nil {
		return
	}
	c.Start()
}
