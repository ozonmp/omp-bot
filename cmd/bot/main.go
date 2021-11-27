package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"

	trv_ticket_api "github.com/ozonmp/trv-ticket-api/pkg/trv-ticket-api"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
	"github.com/ozonmp/omp-bot/internal/app/router"
)

func main() {
	godotenv.Load()

	token := os.Getenv("TOKEN")
	apiAddress := os.Getenv("TRV_TICKET_API_ADDRESS")

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

	conn, err := grpc.Dial(apiAddress, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	router := router.NewRouter(
		context.Background(),
		trv_ticket_api.NewTravelTicketApiServiceClient(conn),
		bot,
	)

	for update := range updates {
		router.HandleUpdate(update)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit
}
