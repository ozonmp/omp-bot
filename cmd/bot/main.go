package main

import (
	"context"
	"fmt"
	cnm_film "github.com/ozonmp/omp-bot/internal/clients/cnm-film"
	"github.com/ozonmp/omp-bot/internal/config"
	"github.com/ozonmp/omp-bot/internal/utils/logger"
	filmApi "github.com/ozonmp/omp-bot/pb/github.com/ozonmp/cnm-film-api/pkg/cnm-film-api"
	"github.com/rs/zerolog/log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
	routerPkg "github.com/ozonmp/omp-bot/internal/app/router"
)

const configPath = "config.yaml"

func main() {
	cfg, err := config.InitConfigYAML(configPath)
	if err != nil {
		log.Fatal().Err(err).Msg("Can't init config")
	}

	ctx := context.Background()
	globLogger, ctx := logger.NewLoggerWithContext(ctx, cfg)

	globLogger.Info().Msg("Creating filmApi client")
	host := fmt.Sprintf("%s:%s", cfg.CnmFilm.Host, cfg.CnmFilm.Port)
	filmApiClient, err := cnm_film.NewClient(ctx, host)
	if err != nil {
		globLogger.Fatal().Err(err).Msg("Failed to create filmApiClient")
	}

	globLogger.Info().Msg("Start loading environment variables")
	if err := godotenv.Load(); err != nil {
		globLogger.Fatal().Err(err).Msg("Failed to load .env file")
	}

	globLogger.Info().Msg("Looking for token")
	token, found := os.LookupEnv("TOKEN")
	if !found {
		globLogger.Fatal().Msg("environment variable TOKEN not found")
	}

	globLogger.Info().Msg("Initializing bot api")
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		globLogger.Fatal().Err(err).Msg("Failed to initialize bot api")
	}

	globLogger.Info().Msg(fmt.Sprintf("Authorized on account %s", bot.Self.UserName))

	u := tgbotapi.UpdateConfig{
		Timeout: 60,
	}

	globLogger.Info().Msg("Start looking for updates")
	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		globLogger.Fatal().Err(err).Msg("Failed to start looking for updates")
	}

	routerHandler := routerPkg.NewRouter(ctx, bot, filmApiClient)

	for update := range updates {
		routerHandler.HandleUpdate(update)
	}
}
