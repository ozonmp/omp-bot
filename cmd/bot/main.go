package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
	"github.com/ozonmp/omp-bot/internal/app/router"
)

func main() {
	godotenv.Load()

	token := os.Getenv("TOKEN")

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	// Uncomment if you want debugging
	// bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.UpdateConfig{
		Timeout: 60,
	}

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Panic(err)
	}

	router := router.NewRouter(bot)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

loop:
	for {
		select {
		case update, ok := <-updates:
			if !ok {
				break loop
			}
			router.HandleUpdate(update)
		case <-quit:
			log.Print("Terminate signal received. Finishing waiting for updates.")
			break loop
		}
	}
}
