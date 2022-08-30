package main

import (
	"context"
	"fmt"
	"github.com/christiansoetanto/better-servant-of-servus-dei/config"
	"github.com/christiansoetanto/better-servant-of-servus-dei/database"
	"github.com/christiansoetanto/better-servant-of-servus-dei/dbot"
	"github.com/christiansoetanto/better-servant-of-servus-dei/provider"
	"github.com/christiansoetanto/better-servant-of-servus-dei/usecase"
	"log"
	"os"
	"os/signal"
	"syscall"
)

const DEVMODE = true

func main() {
	fmt.Println("Hello World!")
	ctx := context.Background()
	cfg := config.Init(DEVMODE)
	database.New(ctx, cfg.AppConfig)
	defer database.Close(ctx)

	providerResource := &provider.Resource{
		AppConfig: cfg.AppConfig,
		Database:  database.GetDBObject(ctx, cfg.AppConfig),
	}

	prov := provider.GetProvider(providerResource)

	session := dbot.New()
	usecaseResource := &usecase.Resource{
		Provider: prov,
		Config:   cfg,
		Session:  session,
	}

	usecaseObject := usecase.GetUsecaseObject(usecaseResource)
	err := usecaseObject.Init(ctx)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		usecaseObject.CloseSessionConnection()
	}()

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Session is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	//syscall.SIGTERM,
	signal.Notify(sc, syscall.SIGINT)
	<-sc

	log.Println("Gracefully shutting down.")
}
