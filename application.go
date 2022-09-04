package main

import (
	"context"
	"fmt"
	"github.com/christiansoetanto/better-servant-of-servus-dei/config"
	"github.com/christiansoetanto/better-servant-of-servus-dei/dbot"
	"github.com/christiansoetanto/better-servant-of-servus-dei/provider"
	"github.com/christiansoetanto/better-servant-of-servus-dei/usecase"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

func main() {
	fmt.Println("Hello World!")
	ctx := context.Background()
	devMode, err := strconv.ParseBool(os.Getenv("DEVMODE"))
	if err != nil {
		log.Fatal("Error parsing DEVMODE environment variable")
		return
	}

	cfg := config.Init(devMode)
	//database.New(ctx, cfg.AppConfig)
	//defer database.Close(ctx)

	providerResource := &provider.Resource{
		AppConfig: cfg.AppConfig,
		//Database:  database.GetDBObject(ctx, cfg.AppConfig),
	}

	prov := provider.GetProvider(providerResource)

	session := dbot.New()
	usecaseResource := &usecase.Resource{
		Provider: prov,
		Config:   cfg,
		Session:  session,
	}

	usecaseObject := usecase.GetUsecaseObject(usecaseResource)
	err = usecaseObject.Init(ctx)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		_ = usecaseObject.CloseSessionConnection()
	}()

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Session is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	//syscall.SIGTERM,
	signal.Notify(sc, syscall.SIGINT)
	<-sc

	log.Println("Gracefully shutting down.")
}
