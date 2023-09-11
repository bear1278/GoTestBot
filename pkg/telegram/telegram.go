package telegram

import (
	"log"

	"github.com/bear1278/GoTestBot/config"
	"github.com/bear1278/GoTestBot/db"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
)



const(
	commandStart="start"
	commandAddNewTask="addnewtask"
	commandDeleteTask="deletetask"
	commandMarkTask="marktask"
	commandListTask="listoftasks"
)


var(
	//ListOfUsers []User
	MapOfTasks=make(map[int64]string)
)



type Bot struct{
	bot *tgbotapi.BotAPI
	db db.DbRepository
	Messages config.Messages
}



func NewBot(bot *tgbotapi.BotAPI, db db.DbRepository, msg config.Messages) *Bot{
	return &Bot{bot: bot, db: db, Messages: msg}
}




func (b *Bot) Start() error{
	log.Printf("Authorized on account %s", b.bot.Self.UserName)

	
	updates,err:=b.initUpdatesChannel()

	if err!=nil{
		return err
	}

	b.handleUpdates(updates)
	return nil
}




func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel){
	for update := range updates {
		if update.Message != nil { // If we got a message
			
			if update.Message.IsCommand(){
				err:=b.handleCommand(update.Message)
				if err!=nil{
					b.handlerError(update.Message.Chat.ID,err)
				}
				continue
			}
			
			err:=b.handleMessege(update.Message)
			if err!=nil{
				b.handlerError(update.Message.Chat.ID, err)
			}
		}
	}
}





func (b *Bot) initUpdatesChannel() (tgbotapi.UpdatesChannel, error){
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60	

	return b.bot.GetUpdatesChan(u)
}