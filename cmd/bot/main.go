package main

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
	routerPkg "github.com/ozonmp/omp-bot/internal/app/router"
	"github.com/ozonmp/omp-bot/internal/config"
	"github.com/ozonmp/omp-bot/internal/pkg/logger"
	"log"
	"os"
)

func main() {
	_ = godotenv.Load()

	token, found := os.LookupEnv("TOKEN")
	if !found {
		log.Panic("environment variable TOKEN not found in .env")
	}
	ctx := context.Background()
	cfg, err := config.ReadConfigYML("config.yml")
	if err != nil {
		logger.FatalKV(ctx, "Failed init configuration", "err", err)
	}

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.UpdateConfig{
		Timeout: 60,
	}

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Panic(err)
	}

	routerHandler := routerPkg.NewRouter(ctx, bot, cfg.Grpc)

	for update := range updates {
		routerHandler.HandleUpdate(update)
	}
}
