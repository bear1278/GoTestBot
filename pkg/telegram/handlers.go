package telegram

import (
	"log"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

func (b *Bot) handleUnknownCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, b.Messages.UnvalidCommand)
	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) handleCommand(command *tgbotapi.Message) error {
	switch command.Command() {
	case commandStart:
		return b.handleStartCommand(command)
	case commandAddNewTask:
		return b.checkUserList(command, "add")
	case commandDeleteTask:
		return b.checkUserList(command, "delete")
	case commandMarkTask:
		return b.checkUserList(command, "mark")
	case commandListTask:
		return b.handleListTaskCommand(command)
	default:
		return b.handleUnknownCommand(command)
	}
	return nil
}

func (b *Bot) handleStartCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, b.Messages.Start)
	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) handleAddTaskCommand(message *tgbotapi.Message) error {

	delete(MapOfTasks, message.Chat.ID)

	check, err := b.db.CheckTask(message.Chat.ID, message.Text)
	if err != nil {
		return nil
	}
	if check {
		return errIndenticalTask
	}

	err = b.db.Add(message.Chat.ID, message.Text)
	if err != nil {
		return err
	}
	msg := tgbotapi.NewMessage(message.Chat.ID, b.Messages.TaskAdded)
	_, err = b.bot.Send(msg)

	return err
}

// исправить  единоообразие методов
func (b *Bot) handleDeleteTaskCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, b.Messages.TaskDelete)

	delete(MapOfTasks, message.Chat.ID)
	check, err := b.db.CheckTask(message.Chat.ID, message.Text)
	if err != nil {
		return nil
	}
	if !check {
		return errTaskNotFound
	}

	err = b.db.Delete(message.Chat.ID, message.Text)
	if err != nil {
		return err
	}
	_, err = b.bot.Send(msg)

	return err
}

func (b *Bot) handleMarkTaskCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, b.Messages.TaskMark)

	delete(MapOfTasks, message.Chat.ID)
	check, err := b.db.CheckTask(message.Chat.ID, message.Text)
	if err != nil {
		return nil
	}
	if !check {
		return errTaskNotFound
	}

	check, err = b.db.CheckAlreadyMark(message.Chat.ID, message.Text)
	if err != nil {
		return nil
	}
	if check {
		return errAlreadyMark
	}

	err = b.db.Mark(message.Chat.ID, message.Text)

	if err != nil {
		return err
	}

	_, err = b.bot.Send(msg)

	return err

}

func (b *Bot) checkUserList(message *tgbotapi.Message, mode string) error {

	MapOfTasks[message.Chat.ID] = mode
	msg := tgbotapi.NewMessage(message.Chat.ID, b.Messages.EnterTask)
	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) handleListTaskCommand(message *tgbotapi.Message) error {

	listOfTasks, err := b.db.List(message.Chat.ID)

	if err != nil {
		return err
	}

	if listOfTasks == "" {
		return errNoTasks
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, listOfTasks)
	_, err = b.bot.Send(msg)
	return err
}

func (b *Bot) handleMessege(message *tgbotapi.Message) error {
	log.Printf("[%s] %s", message.From.UserName, message.Text)
	task, ok := MapOfTasks[message.Chat.ID]
	if ok {
		switch task {
		case "add":
			return b.handleAddTaskCommand(message)
		case "delete":
			return b.handleDeleteTaskCommand(message)
		case "mark":
			return b.handleMarkTaskCommand(message)
		}
	} else {
		return errUnvalidText
	}
	return nil
}
