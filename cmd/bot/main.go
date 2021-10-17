package main

import (
	"log"
	"os"

	"github.com/ozonmp/omp-bot/internal/app/router"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load() // nolint: errcheck

	token, found := os.LookupEnv("TOKEN")
	if !found {
		log.Panic("environment variable TOKEN not found in .env")
	}

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

	r := router.NewRouter(bot)

	for update := range updates {
		r.HandleUpdate(&update)
	}
}
