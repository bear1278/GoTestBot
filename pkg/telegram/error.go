package telegram

import (
	"errors"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

var (
	errIndenticalTask=errors.New("task already exist")
	errNoTasks=errors.New("user hasn't got tasks")
	errUnvalidText=errors.New("unvalid text")
	errTaskNotFound=errors.New("task not found")
	errAlreadyMark=errors.New("task already marked")
)

func (b*Bot) handlerError(ChatID int64, err error){
	msg := tgbotapi.NewMessage(ChatID, b.Messages.UnvalidError)
	switch err{
		case errIndenticalTask:
			msg.Text=b.Messages.IndenticalTask
			b.bot.Send(msg)
		case errNoTasks:
			msg.Text=b.Messages.NoTasks
			b.bot.Send(msg)
		case errUnvalidText:
			msg.Text=b.Messages.UnvalidText
			b.bot.Send(msg)
		case errTaskNotFound:
			msg.Text=b.Messages.TaskNotFound
			b.bot.Send(msg)
		case errAlreadyMark:
			msg.Text=b.Messages.AlreadyMark
			b.bot.Send(msg)
		default:
			b.bot.Send(msg)
	}
}