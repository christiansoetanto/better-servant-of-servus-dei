package main

import (
	"fmt"
	"github.com/christiansoetanto/better-servant-of-servus-dei/config"
	"github.com/christiansoetanto/better-servant-of-servus-dei/dbot"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	fmt.Println("Hello World!")
	cfg := config.Init()

	dbot := dbot.New(cfg)
	err := dbot.NewSession()
	if err != nil {
		log.Fatalf("error creating Discord session: %v", err)
		return
	}

	dbot.LoadAllHandlers()

	err = dbot.OpenConnection()
	if err != nil {
		log.Fatalf("error opening Discord connection: %v", err)
		return
	}

	defer dbot.Session.Close()

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	//syscall.SIGTERM,
	signal.Notify(sc, syscall.SIGINT)
	<-sc

	log.Println("Gracefully shutting down.")
	return
}
