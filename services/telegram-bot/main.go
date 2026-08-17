package main

import (
	"log"
	"os"
	"os/signal"
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)


func helpHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "This is a help message.",
	})
}

func main() {
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	log.Printf("bot token: %v", botToken)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		bot.WithDefaultHandler(helpHandler),
	}

	b, err := bot.New(botToken, opts...)

	b.RegisterHandler(bot.HandlerTypeMessageText, "/help", bot.MatchTypeExact, helpHandler)

	if err != nil {
		log.Fatalf("Failed to start bot: %v", err)
		return
	}

	b.Start(ctx)
}
