package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/router"
	"log"
)

func main() {
	//godotenv.Load()

	token := "2068890040:AAEPqBcrrY0Z967xACYUPeAxhU4Xe0wVoK8"//os.Getenv("2068890040:AAEPqBcrrY0Z967xACYUPeAxhU4Xe0wVoK8")

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

	for update := range updates {
		router.HandleUpdate(update)
	}
}
