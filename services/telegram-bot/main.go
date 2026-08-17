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
	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "This is a help message.",
	})

	if err != nil {
		log.Fatalf("Failed to send help handler message: %v", err)
		return
	}

	log.Print("Sent help message.")
}

func remindHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "What do you want me to remind You in the future?",
	})

	if err != nil {
		log.Fatalf("Failed to send remind handler message: %v", err)
		return
	}

	log.Print("Sent remind message.")
}

func main() {
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		bot.WithDefaultHandler(helpHandler),
	}

	b, err := bot.New(botToken, opts...)

	b.RegisterHandler(bot.HandlerTypeMessageText, "/help", bot.MatchTypeExact, helpHandler)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/remind", bot.MatchTypePrefix, remindHandler)

	if err != nil {
		log.Fatalf("Failed to start bot: %v", err)
		return
	}

	b.Start(ctx)
}
