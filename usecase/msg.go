package usecase

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	TopicFirstMessage = 6
	GroupId           = -1002245243634
)

func (u *UseCase) SendMessageToGroup(msg string) error {
	message := tgbotapi.NewMessage(GroupId, msg)
	message.ReplyToMessageID = TopicFirstMessage
	_, err := u.t.Send(message)
	return err
}
