package usecase

import (
	"github.com/ashurov-imomali/tgbot/config"
	"github.com/ashurov-imomali/tgbot/logger"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type UseCase struct {
	t *tgbotapi.BotAPI
	l logger.ILogger
}

func New(conf config.Configs, l logger.ILogger) IUseCase {
	api, err := tgbotapi.NewBotAPI(conf.Token)
	if err != nil {
		l.Fatal("[GetNewBotApi]", err)
	}
	return &UseCase{
		t: api,
		l: l,
	}
}

type IUseCase interface {
	Pong() string
	SendMessageToGroup(msg string) error
}
