package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/router"
)

func main() {
	config, err := loadConfig()
	if err != nil {
		log.Panicf("load config: %v", err)
	}
	if config.Debug {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
	}

	bot, err := tgbotapi.NewBotAPI(config.Token)
	if err != nil {
		log.Panicf("connect to API: %v", err)
	}
	bot.Debug = config.Debug

	log.Printf("Authorized on account %s", bot.Self.UserName)

	updates, err := bot.GetUpdatesChan(tgbotapi.UpdateConfig{
		Timeout: 60,
	})
	if err != nil {
		log.Panic(err)
	}

	router := router.NewRouter(bot)

	for update := range updates {
		router.HandleUpdate(update)
	}
	log.Printf("Exited after updates channel was closed")
}
