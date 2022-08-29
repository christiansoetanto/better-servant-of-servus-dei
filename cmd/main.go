package main

import (
	"context"
	"fmt"
	"github.com/christiansoetanto/better-servant-of-servus-dei/config"
	"github.com/christiansoetanto/better-servant-of-servus-dei/dbot"
	"github.com/christiansoetanto/better-servant-of-servus-dei/fstore"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	fmt.Println("Hello World!")
	ctx := context.Background()

	cfg := config.Init()
	firestoreClient, err := fstore.Init(ctx)
	defer firestoreClient.Close()

	if err != nil {
		log.Fatalf("Failed to init firestore: %v", err)
		return
	}
	dbot := dbot.New(cfg, firestoreClient)
	err = dbot.NewSession()
	if err != nil {
		log.Fatalf("error creating Discord session: %v", err)
		return
	}

	dbot.LoadAllHandlers()
	dbot.InitAllCronJobs()

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
