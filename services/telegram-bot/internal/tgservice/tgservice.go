package tgservice

import (
	"log"
	"os"
	"os/signal"
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type ConversationState string

const (
	StateDefault ConversationState = ""
	StateWaitingText ConversationState = "waiting text"
)

type SessionState struct {
	LastMessage models.Message
	CurrentState ConversationState
}

type TgService struct {
	b *bot.Bot
	ctx *context.Context
	stop context.CancelFunc

	currentSessionState SessionState
}

func (svc *TgService) defaultHandler() func(ctx context.Context, b *bot.Bot, update *models.Update) {
	return func(ctx context.Context, b *bot.Bot, update *models.Update) {
		switch svc.currentSessionState.CurrentState {
		case StateDefault:
			svc.helpHandler()(ctx, b, update)

		case  StateWaitingText:
			_, err := b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: update.Message.Chat.ID,
				Text:   "Ok, I will remind this message later:",
			})

			if err != nil {
				log.Fatalf("Failed to send help handler message: %v", err)
				return
			}

			_, err = svc.b.ForwardMessage(ctx, &bot.ForwardMessageParams{
				ChatID: update.Message.Chat.ID,
				FromChatID: update.Message.Chat.ID,
				MessageID: update.Message.ID,
			})

			if err != nil {
				log.Fatalf("Failed to send help handler message: %v", err)
				return
			}

			svc.currentSessionState.CurrentState = StateDefault

		default:
			log.Fatalf("Wrong conversation state")
		}
	}
}

func (svc *TgService) helpHandler() func(ctx context.Context, b *bot.Bot, update *models.Update) {
	return func(ctx context.Context, b *bot.Bot, update *models.Update) {
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
}

func (svc *TgService) remindHandler() func(ctx context.Context, b *bot.Bot, update *models.Update) {
	return func(ctx context.Context, b *bot.Bot, update *models.Update) {
		_, err := b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "What do you want me to remind You in the future?",
		})

		if err != nil {
			log.Fatalf("Failed to send remind handler message: %v", err)
			return
		}

		svc.currentSessionState.CurrentState = StateWaitingText

		log.Print("Sent remind message.")
	}
}

func New(token string) *TgService {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)

	svc := &TgService{
		ctx: &ctx,
		stop: stop,
		currentSessionState: SessionState{
			CurrentState: StateDefault,
		},
	}

	opts:= []bot.Option{
		bot.WithDefaultHandler(svc.defaultHandler()),
	}

	b, err := bot.New(token, opts...)
	if err != nil {
		log.Fatalf("Failed to create telegram bot: %v", err)
		return nil
	}

	svc.b = b

	return svc
}

func (svc *TgService) RegisterHandlers() {
	svc.b.RegisterHandler(bot.HandlerTypeMessageText, "/help", bot.MatchTypeExact, svc.helpHandler())
	svc.b.RegisterHandler(bot.HandlerTypeMessageText, "/remind", bot.MatchTypePrefix, svc.remindHandler())
}

func (svc *TgService) Start() {
	defer svc.stop()

	svc.b.Start(*svc.ctx)
}
