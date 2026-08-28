package main

import (
	"os"
	"telegram-bot/internal/tgservice"
)

func main() {
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")

	svc := tgservice.New(botToken)
	if svc == nil {
		return
	}

	svc.RegisterHandlers()
	svc.Start()
}
