package telegram

import (
	"log"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
)


func (b *Bot) handleUnknownCommand(message *tgbotapi.Message) error{
	msg := tgbotapi.NewMessage(message.Chat.ID, b.Messages.UnvalidCommand)
	_,err:=b.bot.Send(msg)
			return err
}


func (b *Bot) handleCommand(command *tgbotapi.Message) error{
	switch command.Command(){
		case commandStart: 
			b.handleStartCommand(command)
		case commandAddNewTask:
			b.checkUserList(command,"add")
		case commandDeleteTask:
			b.checkUserList(command,"delete")
		case commandMarkTask:
			b.checkUserList(command,"mark")
		case commandListTask:
			b.handleListTaskCommand(command)
		default: 
			b.handleUnknownCommand(command)
	}
	return nil
}



func (b *Bot) handleStartCommand(message *tgbotapi.Message) error{
	msg := tgbotapi.NewMessage(message.Chat.ID, b.Messages.Start)
	_,err:=b.bot.Send(msg)
	return err
}


func (b *Bot) handleAddTaskCommand(message *tgbotapi.Message) error{
	// for i,user:=range ListOfUsers{
	// 	if user.CompareId(message.Chat.ID){
	// 		user.SetTask(message.Text)
	// 		ListOfUsers[i]=user
	// 		break
	// 	}
	// }
	
	delete(MapOfTasks,message.Chat.ID)

	check,err:=b.db.CheckTask(message.Chat.ID, message.Text)
	if err!=nil{
		return nil
	}
	if check{
		return errIndenticalTask
	}

	err=b.db.Add(message.Chat.ID,message.Text)
	if err!=nil{
		return err
	}
	msg := tgbotapi.NewMessage(message.Chat.ID, b.Messages.TaskAdded)
	_, err=b.bot.Send(msg)
	
	return err
}

// исправить  единоообразие методов
func (b *Bot) handleDeleteTaskCommand(message *tgbotapi.Message) error{
	msg := tgbotapi.NewMessage(message.Chat.ID, b.Messages.TaskDelete)
	// for i,user:=range ListOfUsers{
	// 	if user.CompareId(message.Chat.ID){
	// 		if !user.DeleteTask(message.Text){
	// 			msg.Text="Task nt found"
	// 		}
	// 		ListOfUsers[i]=user
	// 		break
	// 	}
	// }
	delete(MapOfTasks,message.Chat.ID)
	check,err:=b.db.CheckTask(message.Chat.ID, message.Text)
	if err!=nil{
		return nil
	}
	if !check{
		return errTaskNotFound
	}

	err=b.db.Delete(message.Chat.ID,message.Text)
	if err!=nil{
		return err
	}
	_, err=b.bot.Send(msg)
	
	return err
}

func (b *Bot) handleMarkTaskCommand(message *tgbotapi.Message) error{
	msg := tgbotapi.NewMessage(message.Chat.ID, b.Messages.TaskMark)
	// for i,user:=range ListOfUsers{
	// 	if user.CompareId(message.Chat.ID){
	// 		if !user.MarkTask(message.Text){
	// 			msg.Text="Task not found"
	// 		}
	// 		ListOfUsers[i]=user
	// 		break
	// 	}
	// }
	delete(MapOfTasks,message.Chat.ID)
	check,err:=b.db.CheckTask(message.Chat.ID, message.Text)
	if err!=nil{
		return nil
	}
	if !check{
		return errTaskNotFound
	}

	check,err=b.db.CheckAlreadyMark(message.Chat.ID, message.Text)
	if err!=nil{
		return nil
	}
	if check{
		return errAlreadyMark
	}

	err=b.db.Mark(message.Chat.ID,message.Text)

	if err!=nil{
		return err
	}
	
	_, err=b.bot.Send(msg)
	
	return err

}


func (b *Bot) checkUserList (message *tgbotapi.Message, mode string) error{
	// if len(ListOfUsers)>0{
	// 	count:=0
	// 	for _,v:=range ListOfUsers{
			
	// 		if v.CompareId(message.Chat.ID){
	// 			break
	// 		}	
	// 		count ++
	// 	}
	// 	if count==len(ListOfUsers) && !ListOfUsers[len(ListOfUsers)-1].CompareId(message.Chat.ID){
	// 		ListOfUsers=append(ListOfUsers, *NewUser(message.Chat.ID))
	// 	}
	// }else{
	// 	ListOfUsers=append(ListOfUsers, *NewUser(message.Chat.ID))
	// }
	MapOfTasks[message.Chat.ID]=mode
	msg := tgbotapi.NewMessage(message.Chat.ID, b.Messages.EnterTask)
	_,err:=b.bot.Send(msg)
	return err
}


func (b *Bot) handleListTaskCommand(message *tgbotapi.Message) error{
	// user:=NewUser(message.Chat.ID)
	// for _,i:=range ListOfUsers{
	// 	if i.CompareId(message.Chat.ID){
	// 		user=&i
	// 		break
	// 	}
	// }
	listOfTasks, err:=b.db.List(message.Chat.ID)

	if err!=nil{
		return err
	}

	if listOfTasks==""{
		return errNoTasks
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, listOfTasks)
	_,err=b.bot.Send(msg)
	return err
}


func (b *Bot) handleMessege(message *tgbotapi.Message) error{
	log.Printf("[%s] %s", message.From.UserName, message.Text)
	task,ok:=MapOfTasks[message.Chat.ID]
	if ok{
		switch task{
		case "add":
			return b.handleAddTaskCommand(message)
		case "delete":
			return b.handleDeleteTaskCommand(message)
		case "mark":
			return b.handleMarkTaskCommand(message)
		}
		}else{
			return errUnvalidText
		}
		return nil
	}