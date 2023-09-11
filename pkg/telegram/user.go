package telegram

import (
	"log"
	"strconv"
	"time"
)
type User struct{
	ChatID int64
	Tasks []Task
}

func  NewUser(id int64) *User {
	task:=make([]Task,0)
	return &User{ChatID: id, Tasks: task}
}




func (u *User) GetListTask() string{
	list:=""
	integer:=""
	
	for i,v:=range u.Tasks{
		
		integer=strconv.Itoa(i+1)
		list+=integer+". "+v.GetDataOfTask()+"\n"
		
	}
	
	return list
}

func (u *User)SetTask(task string, t time.Time, mark string) {
	u.Tasks=append(u.Tasks, *newTask(task,t, mark))
	
	log.Printf("%s   %s",u.Tasks[0],"set task")
	
}


