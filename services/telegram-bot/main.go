package main

import (
	"log"
	"os"
	"os/signal"
	"context"

	"github.com/joho/godotenv"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func handler(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   update.Message.Text,
	})
}

func main() {
	err:= godotenv.Load()
	if err != nil {
		log.Fatalf("Failed to load env: %v", err)
		return
	}

	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	log.Printf("bot token: %v", botToken)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		bot.WithDefaultHandler(handler),
	}

	b, err := bot.New(botToken, opts...)

	if err != nil {
		log.Fatalf("Failed to start bot: %v", err)
		return
	}

	b.Start(ctx)
}
